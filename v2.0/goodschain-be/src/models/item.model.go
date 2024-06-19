package models

type Item struct {
	Id                string
	CreationDate      string
	RemainingQuantity int
	TotalQuantity     int
	ProductType       string
	ExpirationDate    string
}
