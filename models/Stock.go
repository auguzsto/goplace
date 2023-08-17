package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model `gorm:"serializer:json"`
	ProductID  int     `gorm:"serializer:json"`
	Product    Product `gorm:"serializer:json"`
	Amount     float64 `gorm:"serializer:json"`
}
