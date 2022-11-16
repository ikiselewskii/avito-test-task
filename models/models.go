package models

import (
	"github.com/uptrace/bun"
	"time"
)

type MyBaseModel struct {
	bun.BaseModel

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Customer struct {
	ID          int `bun:"id,pk" json:"id" binding:"required"`
	MyBaseModel `bun:"table:customers,alias:c"`
	Balance     int `json:"balance" binding:"required"`
}
type Reservation struct {
	ID          int `bun:"id,pk,autoincrement"`
	MyBaseModel `bun:"table:reservations,alias:r"`
	CustomerID      int
	Customer        Customer `bun:"rel:belongs-to,join:customer_id=id"`
	ProductID   string
	OrderID     string
	Price       string
	Status      string
}
