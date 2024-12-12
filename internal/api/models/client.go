package models

type Client struct {
	ID int `Json:"id"`
	Entity string `Json:"entity"`
	City string `Json:"city"`
	Uf string `Json:"uf"`
	Tel string `Json:"tel"`
	Email string `Json:"email"`
}