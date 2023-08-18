package dto

type ProductDTO struct {
	Id    int     `json:"id"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}
