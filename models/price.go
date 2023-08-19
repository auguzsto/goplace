package models

type Price struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	Product      Product
	ProductID    uint
	QuantityMost float64
	Price        float64
	PriceMost    float64
}
