package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"advertising-task/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(router *echo.Echo) {
	router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
