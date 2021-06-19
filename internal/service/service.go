package service

import (
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository"
)

type Advert interface {
	GetAll(page int, sort string) ([]*model.Advert, error)
	GetById(advertId int, fields bool) (*model.DetailedAdvert, error)
	Create(advert *model.DetailedAdvert) (int, error)
}

type Service struct {
	Advert
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Advert: NewAdvertService(repos.Advert),
	}
}
