package service

import (
	"time"

	. "github.com/VolkovEgor/advertising-task/internal/error"
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

func (s *AdvertService) GetAll(page int, sortParams *model.SortParams) ([]*model.Advert, error) {
	return s.repo.GetAll(page, sortParams)
}

func (s *AdvertService) GetById(advertId int, fields bool) (*model.DetailedAdvert, error) {
	return s.repo.GetById(advertId, fields)
}

func (s *AdvertService) Create(advert *model.DetailedAdvert) (int, error) {

	if advert.Title == "" || len(advert.Title) > 200 {
		return 0, ErrWrongTitle
	}

	if advert.Description == "" || len(advert.Description) > 1000 {
		return 0, ErrWrongDescription
	}

	if advert.Photos == nil || len(advert.Photos) > 3 {
		return 0, ErrWrongPhotos
	}

	if advert.Price < 0 {
		return 0, ErrNotPositivePrice
	}

	advert.CreationDate = time.Now().Unix()

	advertId, err := s.repo.Create(advert)
	if err != nil {
		return 0, err
	}

	return advertId, nil
}
