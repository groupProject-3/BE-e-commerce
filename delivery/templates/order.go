package templates

import (
	"time"
)

type OrderRequest struct {
	Payment_method_id uint   `json:"payment_method_id"`
	Status            string `json:"status"`
}

type OrderResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Payment_method_id uint                  `json:"payment_method_id"`
	Status            string                `json:"status"`
	OrderDetails      []OrderDetailResponse `json:"orderDetails"`
}
