package models


type Book struct {
	ID string `json:"id"`
	Title string	`json:"title"`
	Author string	`json:"author"`
	ISBN string		`json:"isbn"`
	PublishDate string	`json:"publish_date"`
	
}