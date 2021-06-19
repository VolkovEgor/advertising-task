package postgres

import (
	"github.com/VolkovEgor/advertising-task/internal/model"
	"github.com/lib/pq"

	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	NumElementsInPage = 10
)

type AdvertPg struct {
	db *sqlx.DB
}

func NewAdvertPg(db *sqlx.DB) *AdvertPg {
	return &AdvertPg{db: db}
}

func (r *AdvertPg) GetAll(page int, sortParams *model.SortParams) ([]*model.Advert, error) {
	var adverts []*model.Advert
	var query string
	start := (page - 1) * NumElementsInPage

	if sortParams != nil {
		sortParams.Field = "a." + sortParams.Field
		query = fmt.Sprintf(`SELECT a.id, a.title, a.photos[1], a.price
			FROM %s AS a
			ORDER BY %s %s
			LIMIT %d OFFSET %d`, advertsTable, sortParams.Field, sortParams.Order,
			NumElementsInPage, start)
	} else {
		query = fmt.Sprintf(`SELECT a.id, a.title, a.photos[1], a.price
			FROM %s AS a
			LIMIT %d OFFSET %d`, advertsTable, NumElementsInPage, start)
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		advert := &model.Advert{}
		err := rows.Scan(&advert.Id, &advert.Title, &advert.MainPhoto, &advert.Price)

		if err != nil {
			return nil, err
		}

		adverts = append(adverts, advert)
	}

	return adverts, err
}

func (r *AdvertPg) GetById(advertId int, fields bool) (*model.DetailedAdvert, error) {
	var err error
	advert := &model.DetailedAdvert{}

	if fields {
		query := fmt.Sprintf(`SELECT a.id, a.title, a.description, a.photos, a.price
		FROM %s AS a
		WHERE a.id = $1`, advertsTable)

		row := r.db.QueryRow(query, advertId)
		err = row.Scan(&advert.Id, &advert.Title, &advert.Description,
			pq.Array(&advert.Photos), &advert.Price)
	} else {
		advert.Photos = make([]string, 1)
		query := fmt.Sprintf(`SELECT a.id, a.title, a.photos[1], a.price
		FROM %s AS a
		WHERE a.id = $1`, advertsTable)

		row := r.db.QueryRow(query, advertId)
		err = row.Scan(&advert.Id, &advert.Title, &advert.Photos[0], &advert.Price)
	}

	return advert, err
}
