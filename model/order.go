package model

import "time"

type Order struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	ProductID  int64     `json:"product_id" db:"product_id"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
