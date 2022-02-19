package templates

import (
	"time"
)

type OrderRequest struct {
	Payment_method_id uint   `json:"payment_method_id"`
	Status            string `json:"status"`
	PhoneNumber       uint   `json:"phonenumber"`
}

type OrderResponse struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	Name              string         `json:"name"`
	Payment_method_id uint           `json:"payment_method_id"`
	Status            string         `json:"status"`
	PhoneNumber       uint           `json:"phonenumber"`
	OrderDetails      []CartResponse `json:"orderDetails"`
}
