package service

import (
	"strings"
	"time"

	. "github.com/VolkovEgor/advertising-task/internal/error"
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository"
)

const (
	sortPriceDesc = "price_desc"
	sortPriceAsc  = "price_asc"
	sortDateDesc  = "date_desc"
	sortDateAsc   = "date_asc"
)

type Map map[string]interface{}

type AdvertService struct {
	repo repository.Advert
}

func NewAdvertService(repo repository.Advert) *AdvertService {
	return &AdvertService{repo: repo}
}

func (s *AdvertService) GetAll(page int, sort string) ([]*model.Advert, error) {
	var sortField, order string

	if sort != "" {
		switch sort {
		case sortPriceDesc, sortPriceAsc:
			sortField = "price"
			order = parseOrder(sort)

		case sortDateDesc, sortDateAsc:
			sortField = "creation_date"
			order = parseOrder(sort)

		default:
			return nil, ErrWrongSortParams
		}
	}

	adverts, err := s.repo.GetAll(page, sortField, order)
	if err != nil {
		return nil, err
	}

	if adverts == nil {
		return nil, ErrWrongPageNumber
	}

	return adverts, nil
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

	return s.repo.Create(advert)
}

func parseOrder(sort string) string {
	slice := strings.Split(sort, "_")
	return slice[len(slice)-1]
}
