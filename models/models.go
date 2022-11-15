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

type User struct {
	ID          int `bun:"id,pk"`
	MyBaseModel `bun:"table:users,alias:u"`
	Balance     int
}
type Reservation struct {
	ID          int `bun:"id,pk,autoincrement"`
	MyBaseModel `bun:"table:reservations,alias:r"`
	UserID      int
	User        User `bun:"rel:belongs-to,join:user_id=id"`
	ProductID   string
	OrderID     string
	Price       string
	Status      string
}
