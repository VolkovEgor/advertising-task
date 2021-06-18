package repository

import (
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/VolkovEgor/advertising-task/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Advert interface {
	GetAll(page int) ([]*model.Advert, error)
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Advert: postgres.NewAdvertPg(db),
	}
}
