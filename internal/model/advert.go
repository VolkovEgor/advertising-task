package model

type Advert struct {
	Id           int      `json:"id" bd:"id"`
	Title        string   `json:"title" bd:"title" valid:"length(1|200)"`
	Description  string   `json:"description" bd:"description" valid:"length(1|1000)"`
	Photos       []string `json:"photos" bd:"photos"`
	Price        int      `json:"price" bd:"price" valid:"type(int)"`
	CreationDate int64    `json:"creation_date" bd:"creation_date" valid:"type(int)"`
}
