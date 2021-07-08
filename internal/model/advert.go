package model

type Advert struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	MainPhoto string `json:"main_photo" db:"main_photo"`
	Price     int    `json:"price" db:"price"`
}

type DetailedAdvert struct {
	Id           int      `json:"id" db:"id"`
	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description,omitempty" db:"description"`
	Photos       []string `json:"photos" db:"photos"`
	Price        int      `json:"price" db:"price"`
	CreationDate int64    `json:"creation_date,omitempty" db:"creation_date"`
}
