package handler

import (
	"net/http"
	"strconv"

	"github.com/VolkovEgor/advertising-task/internal/model"

	"github.com/labstack/echo/v4"
)

const (
	priceFieldName        = "price"
	creationDateFieldName = "creation_date"
	descOrder             = "desc"
	ascOrder              = "asc"
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
// @Param sort query string false "Sort field"
// @Param order query string false "Order"
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

	field := ctx.QueryParam("sort")
	order := ctx.QueryParam("order")
	var sortParams *model.SortParams

	if field != "" {
		if field != priceFieldName && field != creationDateFieldName || order != descOrder && order != ascOrder {
			response.Error(http.StatusBadRequest, "Invalid sort params")
			return SendError(ctx, response)
		} else {
			sortParams = &model.SortParams{
				Field: field,
				Order: order,
			}
		}
	}

	response = h.services.Advert.GetAll(page, sortParams)
	return ctx.JSON(http.StatusOK, response.Data)
}
