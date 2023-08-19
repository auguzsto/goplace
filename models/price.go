package models

type Price struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	Product      Product `json:"product"`
	ProductID    uint    `json:"-" gorm:"unique"`
	QuantityMost float64 `json:"quantity_most"`
	Price        float64 `json:"price"`
	PriceMost    float64 `json:"price_most"`
	PriceOffer   float64 `json:"price_offer"`
}
