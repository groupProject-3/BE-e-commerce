package templates

import "time"

type OrderDetailRequest struct {
	Cart_id  uint `json:"cart_id"`
	Order_id uint `json:"order_id"`
	Qty      uint `json:"qty"`
	Price    uint `json:"price"`
}

type OrderDetailResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Cart_id  uint `json:"cart_id"`
	Order_id uint `json:"order_id"`
	Qty      uint `json:"qty"`
	Price    uint `json:"price"`
}
