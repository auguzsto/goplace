package models

type Price struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	Product      Product
	ProductID    uint    `json:"product_id" gorm:"unique"`
	ProductCode  string  `json:"product_code" gorm:"unique"`
	QuantityMost float64 `json:"quantity_most"`
	Price        float64 `json:"price"`
	PriceMost    float64 `json:"price_most"`
	PriceOffer   float64 `json:"price_offer"`
}
