package service

import (
	"net/http"
	"time"

	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository"
)

type Map map[string]interface{}

type AdvertService struct {
	repo repository.Advert
}

func NewAdvertService(repo repository.Advert) *AdvertService {
	return &AdvertService{repo: repo}
}

func (s *AdvertService) GetAll(page int, sortParams *model.SortParams) *model.ApiResponse {
	r := &model.ApiResponse{}
	adverts, err := s.repo.GetAll(page, sortParams)
	if err != nil {
		r.Error(http.StatusInternalServerError, err.Error())
		return r
	}

	r.Set(http.StatusOK, "OK", Map{"adverts": adverts})
	return r
}

func (s *AdvertService) GetById(advertId int, fields bool) *model.ApiResponse {
	r := &model.ApiResponse{}

	advert, err := s.repo.GetById(advertId, fields)
	if err != nil {
		r.Error(http.StatusInternalServerError, err.Error())
		return r
	}

	r.Set(http.StatusOK, "OK", Map{"advert": advert})
	return r
}

func (s *AdvertService) Create(advert *model.DetailedAdvert) *model.ApiResponse {
	r := &model.ApiResponse{}

	if len(advert.Title) > 200 || len(advert.Description) > 1000 || len(advert.Photos) > 3 || advert.Price < 0 {
		r.Error(http.StatusBadRequest, "Invalid input data")
		return r
	}

	advert.CreationDate = time.Now().Unix()

	advertId, err := s.repo.Create(advert)
	if err != nil {
		r.Error(http.StatusInternalServerError, err.Error())
		return r
	}

	r.Set(http.StatusOK, "OK", Map{"advertId": advertId})
	return r
}
