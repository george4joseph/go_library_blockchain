package models


type BookIssue struct { 

	BookID string `json:"book_id"`
	User string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis bool `json:"is_genesis"`
}
