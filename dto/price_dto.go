package dto

type PriceDTO struct {
	Id           uint    `json:"id"`
	ProductID    uint    `json:"product_id"`
	ProductCode  string  `json:"product_code"`
	QuantityMost float64 `json:"quantity_most"`
	Price        float64 `json:"price"`
	PriceMost    float64 `json:"price_most"`
	PriceOffer   float64 `json:"price_offer"`
}
