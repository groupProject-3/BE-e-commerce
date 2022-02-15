package product

import "time"

type ProductRequest struct {
	Name            string `json:"name"`
	Product_type_id uint   `json:"product_type_id"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Description     string `json:"description"`
}

type ProductResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name            string `json:"name"`
	Product_type_id uint   `json:"product_type_id"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
	Description     string `json:"description"`
}
