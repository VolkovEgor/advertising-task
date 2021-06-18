package repository

import (
	"github.com/jmoiron/sqlx"
)

type Advert interface {
}

type Repository struct {
	Advert
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
