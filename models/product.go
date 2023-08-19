package models

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code"`
	Quantity  float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
