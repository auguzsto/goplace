package dto

type ProductDTO struct {
	Id        int     `json:"id"`
	Code      string  `json:"code"`
	Quantity  float64 `json:"quantity"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
}
