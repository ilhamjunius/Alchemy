package models

type Book struct {
	ID     string `json:"ID" form:"ID"`
	Tittle string `json:"tittle" form:"tittle"`
	Author string `json:"author" form:"author"`
}
