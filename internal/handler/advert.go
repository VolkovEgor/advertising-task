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
		restaurants.GET("/:aid", h.getAdvertById)
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

	if field != "" && order != "" {
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

// @Summary Get Advert By Id
// @Tags adverts
// @Description Get advert by id
// @ModuleID getAdvertById
// @Accept  json
// @Produce  json
// @Param aid path string true "Advert id"
// @Param fields query string false "Fields"
// @Success 200 {object} model.DetailedAdvert
// @Failure 400 {object} response
// @Failure 404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /adverts/{aid} [get]
func (h *Handler) getAdvertById(ctx echo.Context) error {
	response := &model.ApiResponse{}

	advertId, err := strconv.Atoi(ctx.Param("aid"))
	if err != nil || advertId == 0 {
		response.Error(http.StatusBadRequest, "Invalid advertId")
		return SendError(ctx, response)
	}

	var boolFields bool
	fields := ctx.QueryParam("fields")
	if fields != "" {
		boolFields, err = strconv.ParseBool(fields)
		if err != nil {
			response.Error(http.StatusBadRequest, "Invalid fields parameter")
			return SendError(ctx, response)
		}
	}

	response = h.services.Advert.GetById(advertId, boolFields)
	return ctx.JSON(http.StatusOK, response.Data)
}
