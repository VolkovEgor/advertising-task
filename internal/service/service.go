package service

import "advertising-task/internal/repository"

type Advert interface {
}

type Service struct {
	Advert
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
