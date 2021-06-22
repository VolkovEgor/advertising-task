package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VolkovEgor/advertising-task/internal/handler"
	"github.com/VolkovEgor/advertising-task/internal/repository"
	"github.com/VolkovEgor/advertising-task/internal/repository/postgres"
	"github.com/VolkovEgor/advertising-task/internal/service"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	dbPrefix := viper.GetString("db.name") + "."

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString(dbPrefix + "host"),
		Port:     viper.GetString(dbPrefix + "port"),
		Username: viper.GetString(dbPrefix + "username"),
		DBName:   viper.GetString(dbPrefix + "dbname"),
		SSLMode:  viper.GetString(dbPrefix + "sslmode"),
		Password: viper.GetString(dbPrefix + "password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := echo.New()
	app.Use(middleware.Logger())
	handlers.Init(app)

	go func() {
		if err := app.Start(viper.GetString("port")); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("failed to listen: %s", err.Error())
		}
	}()

	log.Println("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Gracefully shutting down...")
	if err := app.Shutdown(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
