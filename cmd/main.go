package main

import (
	"log"

	"github.com/VolkovEgor/advertising-task/internal/handler"
	"github.com/VolkovEgor/advertising-task/internal/repository"
	"github.com/VolkovEgor/advertising-task/internal/repository/postgres"
	"github.com/VolkovEgor/advertising-task/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := echo.New()
	app.Use(middleware.Logger())
	handlers.Init(app)

	if err := app.Start(viper.GetString("port")); err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
