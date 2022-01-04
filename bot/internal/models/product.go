package models

type Product struct {
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
}


