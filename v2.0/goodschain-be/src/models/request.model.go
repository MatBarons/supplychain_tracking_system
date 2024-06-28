package models

type Request struct {
	Id          string
	User        string
	ItemID      string
	IsConfirmed bool
	Borrower    string
}
