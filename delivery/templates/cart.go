package templates

import "time"

type CartRequest struct {
	Product_id uint   `json:"product_id"`
	Qty        uint   `json:"qty"`
	Status     string `json:"status"`
}

type CartResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product_id  uint   `json:"product_id"`
	Product_qty int   `json:"product_qty"`
	PriceTotal  uint   `json:"pricetotal"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Qty         uint   `json:"qty"`
	Price       uint    `json:"price"`
	Status      string `json:"status"`
}
