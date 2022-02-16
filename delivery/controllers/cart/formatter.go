package cart

import "time"

type CartRequest struct {
	Product_id uint   `json:"product_id"`
	Qty        uint   `json:"qty"`
	Status     string `json:"status"`
}

type CartResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product_id   uint   `json:"product_id"`
	Qty          uint   `json:"qty"`
	Status       string `json:"status"`
	Product_name string `json:"product_name"`
}
