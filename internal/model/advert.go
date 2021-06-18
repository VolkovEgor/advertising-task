package model

type Advert struct {
	Id        int    `json:"id" bd:"id"`
	Title     string `json:"title" bd:"title"`
	MainPhoto string `json:"main_photo" bd:"main_photo"`
	Price     int    `json:"price" bd:"price"`
}

type DetailedAdvert struct {
	Id           int      `json:"id" bd:"id"`
	Title        string   `json:"title" bd:"title"`
	Description  string   `json:"description" bd:"description"`
	Photos       []string `json:"photos" bd:"photos"`
	Price        int      `json:"price" bd:"price"`
	CreationDate int64    `json:"creation_date" bd:"creation_date"`
}
