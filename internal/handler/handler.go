package handler

import (
	"net/http"

	_ "github.com/VolkovEgor/advertising-task/docs/swagger"
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/service"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type response struct {
	Message string `json:"message"`
}

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

	api := router.Group("/api")
	{
		h.initAdvertRoutes(api)
	}
}

func SendError(ctx echo.Context, r *model.ApiResponse) error {
	return ctx.JSON(r.Code, echo.Map{"error": r.Message})
}
