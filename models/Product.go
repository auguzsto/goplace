package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string  `gorm:"serializer:json"`
	Price float64 `gorm:"serializer:json"`
}
