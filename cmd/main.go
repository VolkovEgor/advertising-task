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

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
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
