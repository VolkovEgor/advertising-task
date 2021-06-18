package handler

import (
	_ "advertising-task/docs/swagger"
	"advertising-task/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// @title Advertising Task
// @version 1.0
// @description API Server for Advertising Task

// @host localhost:9000
// @BasePath /api/

func (h *Handler) Init(router *echo.Echo) {
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
