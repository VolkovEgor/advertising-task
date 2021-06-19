package repository

import (
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Advert interface {
	GetAll(page int, sortParams *model.SortParams) ([]*model.Advert, error)
	GetById(advertId int, fields bool) (*model.DetailedAdvert, error)
	Create(advert *model.DetailedAdvert) (int, error)
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advert: postgres.NewAdvertPg(db),
	}
}
