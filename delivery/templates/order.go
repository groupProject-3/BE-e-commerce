package templates

import (
	"time"
)

type OrderRequest struct {
	Payment_method_id uint   `json:"payment_method_id"`
	Status            string `json:"status"`
	PhoneNumber       uint   `json:"phonenumber"`
	Address           string `json:"address"`
}

type OrderResponse struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	Name              string         `json:"name"`
	Payment_method_id uint           `json:"payment_method_id"`
	Status            string         `json:"status"`
	Address           string         `json:"address"`
	PhoneNumber       uint           `json:"phonenumber"`
	OrderDetails      []CartResponse `json:"orderDetails"`
}
