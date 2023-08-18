package models

type Stock struct {
	Id        uint
	ProductID uint
	Product   Product
	Amount    float64
}
