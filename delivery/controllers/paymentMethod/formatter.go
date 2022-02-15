package paymentmethod

import "time"

type PaymentMethodRequest struct {
	Name string `json:"name"`
}

type PaymentMethodResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string `json:"name"`
}
