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
	ID          int `bun:"id,pk" json:"id"`
	MyBaseModel `bun:"table:customers,alias:c"`
	Balance     int `json:"balance"`
}
type Transaction struct {
	ID          int `bun:"id,pk,autoincrement" json:"-"`
	MyBaseModel `bun:"table:transactions,alias:t"`
	FromID      int      `json:"from_id" binding:"required"`
	From        Customer `bun:"rel:belongs-to,join:from_id=id" json:"-"`
	ToID        int      `json:"to_id,omitempty"`
	To          Customer `bun:"rel:belongs-to,join:to_id=id" json:"-"`
	ProductID   int   `json:"product_id,omitempty"`
	OrderID     int   `json:"order_id,omitempty"`
	Amount      int      `json:"amount" binding:"required"`
	Type        int16    `json:"-"` //0 - purchase goods, 1 - replenish user balance, 2 - user to user transfer
	Status      string   `json:"-"`
}
