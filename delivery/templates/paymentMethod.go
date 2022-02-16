package templates

import "time"

type PaymentMethodRequest struct {
	Name string `json:"name"`
}

type PaymentMethodResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`
}
