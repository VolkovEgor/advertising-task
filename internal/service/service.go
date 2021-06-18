package service

import (
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository"
)

type Advert interface {
	GetAll(page int, sortParams *model.SortParams) *model.ApiResponse
}

type Service struct {
	Advert
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Advert: NewAdvertService(repos.Advert),
	}
}
