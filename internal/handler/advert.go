package handler

import (
	"net/http"
	"strconv"

	. "github.com/VolkovEgor/advertising-task/internal/error"
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/asaskevich/govalidator"

	"github.com/labstack/echo/v4"
)

func (h *Handler) initAdvertRoutes(api *echo.Group) {
	adverts := api.Group("/adverts")
	{
		adverts.GET("", h.getAdverts)
		adverts.GET("/:aid", h.getAdvertById)
		adverts.POST("", h.createAdvert)
	}
}

// @Summary Get All Adverts
// @Tags adverts
// @Description Get all adverts
// @ModuleID getAllAdverts
// @Accept  json
// @Produce  json
// @Param page query string true "Page"
// @Param sort query string false "Sort field and order"
// @Success 200 {array} model.Advert
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /adverts [get]
func (h *Handler) getAdverts(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		return SendError(ctx, http.StatusBadRequest, ErrWrongPageNumber)
	}

	sort := ctx.QueryParam("sort")
	adverts, err := h.services.Advert.GetAll(page, sort)
	if err != nil {
		if err == ErrWrongSortParams {
			return SendError(ctx, http.StatusBadRequest, err)
		}
		return SendError(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, adverts)
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
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /adverts/{aid} [get]
func (h *Handler) getAdvertById(ctx echo.Context) error {

	advertId, err := strconv.Atoi(ctx.Param("aid"))
	if err != nil || advertId <= 0 {
		return SendError(ctx, http.StatusBadRequest, ErrWrongAdvertId)
	}

	var boolFields bool
	fields := ctx.QueryParam("fields")
	if fields != "" {
		boolFields, err = strconv.ParseBool(fields)
		if err != nil {
			return SendError(ctx, http.StatusBadRequest, ErrWrongFieldsParam)
		}
	}

	advert, err := h.services.Advert.GetById(advertId, boolFields)
	if err != nil {
		return SendError(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, advert)
}

type advertInput struct {
	Title       string   `json:"title" valid:"length(1|200)"`
	Description string   `json:"description" valid:"length(1|1000)"`
	Photos      []string `json:"photos"`
	Price       int      `json:"price" valid:"type(int)"`
}

// @Summary Create Advert
// @Tags adverts
// @Description Create advert
// @ModuleID createAdvert
// @Accept  json
// @Produce  json
// @Param input body advertInput true "advert input"
// @Success 200 {object} advertIdResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /adverts [post]
func (h *Handler) createAdvert(ctx echo.Context) error {
	var input advertInput

	if err := ctx.Bind(&input); err != nil {
		return SendError(ctx, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(input); err != nil {
		return SendError(ctx, http.StatusBadRequest, err)
	}

	advert := &model.DetailedAdvert{
		Title:       input.Title,
		Description: input.Description,
		Photos:      input.Photos,
		Price:       input.Price,
	}

	advertId, err := h.services.Advert.Create(advert)
	if err != nil {
		if err == ErrWrongTitle || err == ErrWrongDescription ||
			err == ErrWrongPhotos || err == ErrNotPositivePrice {
			return SendError(ctx, http.StatusBadRequest, err)
		}
		return SendError(ctx, http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, advertIdResponse{advertId})
}
