package templates

import "time"

type OrderDetailRequest struct {
	Cart_id  uint `json:"cart_id"`
	Order_id uint `json:"order_id"`
}

type OrderDetailResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name   string `json:"name"`
	Image  string `json:"image"`
	Qty    uint   `json:"qty"`
	Price  int    `json:"price"`
	Status string `json:"status"`
}
