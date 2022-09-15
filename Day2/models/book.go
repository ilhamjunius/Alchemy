package models

type Book struct {
	ID     int    `json:"id" form:"id"`
	Tittle string `json:"tittle" form:"tittle"`
	Author string `json:"author" form:"author"`
}
