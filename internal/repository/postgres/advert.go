package postgres

import (
	"strings"

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

func (r *AdvertPg) GetAll(page int, sortField, order string) ([]*model.Advert, error) {
	var adverts []*model.Advert
	var query string
	start := (page - 1) * NumElementsInPage

	if sortField != "" && order != "" {
		sortField = "a." + sortField
		query = fmt.Sprintf(`SELECT a.id, a.title, a.photos[1], a.price
			FROM %s AS a
			ORDER BY %s %s
			LIMIT %d OFFSET %d`, advertsTable, sortField, order,
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

func (r *AdvertPg) Create(advert *model.DetailedAdvert) (int, error) {
	var advertId int
	photos := strings.Join(advert.Photos, `", "`)
	photos = `{"` + photos + `"}`

	query := fmt.Sprintf(
		`INSERT INTO %s
		(title, description, photos, price, creation_date)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`, advertsTable)

	row := r.db.QueryRow(query, advert.Title, advert.Description, photos, advert.Price, advert.CreationDate)
	if err := row.Scan(&advertId); err != nil {
		return 0, err
	}

	return advertId, nil
}
