package templates

import "time"

type ProductRequest struct {
	Name            string `json:"name"`
	Image           string `json:"image"`
	Product_type_id uint   `json:"product_type_id"`
	Price           uint    `json:"price"`
	Qty             uint    `json:"qty"`
	Description     string `json:"description"`
}

type ProductResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name              string `json:"name"`
	Image             string `json:"image"`
	Price             uint    `json:"price"`
	Qty               uint    `json:"qty"`
	Description       string `json:"description"`
	Product_type_name string `json:"product_type_name"`
}
