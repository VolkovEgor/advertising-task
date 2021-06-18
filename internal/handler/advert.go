package handler

import (
	"net/http"
	"strconv"

	"github.com/VolkovEgor/advertising-task/internal/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) initAdvertRoutes(api *echo.Group) {
	restaurants := api.Group("/adverts")
	{
		restaurants.GET("", h.getAdverts)
	}
}

// @Summary Get All Adverts
// @Tags adverts
// @Description Get all adverts
// @ModuleID getAllAdverts
// @Accept  json
// @Produce  json
// @Param page query string true "Page"
// @Success 200 {array} model.Advert
// @Failure 400 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /adverts [get]
func (h *Handler) getAdverts(ctx echo.Context) error {
	response := &model.ApiResponse{}
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page == 0 {
		response.Error(http.StatusBadRequest, "Invalid page number")
		return SendError(ctx, response)
	}

	response = h.services.Advert.GetAll(page)
	return ctx.JSON(http.StatusOK, response.Data)
}
