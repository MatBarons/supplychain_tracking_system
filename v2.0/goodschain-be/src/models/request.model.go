package models

type Request struct {
	Id          string
	UserID      string
	ItemID      string
	IsConfirmed bool
	Requester   string
}
