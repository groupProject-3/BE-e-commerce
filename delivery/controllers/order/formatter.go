package order

import (
	"be/delivery/controllers/orderDetail"
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

	Payment_method_id uint                              `json:"payment_method_id"`
	Total_qty         uint                              `json:"total_qty"`
	Total_price       uint                              `json:"total_price"`
	Status            string                            `json:"status"`
	OrderDetails      []orderDetail.OrderDetailResponse `json:"orderDetails"`
}
