package models

import (
	paymentmethod "be/delivery/controllers/paymentMethod"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model

	Name   string
	Orders []Order `gorm:"foreignKey:Payment_method_id"`
}

func (p *PaymentMethod) ToPaymentMethodResponse() paymentmethod.PaymentMethodResponse{
	return paymentmethod.PaymentMethodResponse{
		ID: p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,

		Name: p.Name,
	}
}