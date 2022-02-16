package paymentmethod

import "be/models"

type PaymentMethod interface {
	Create(pm models.PaymentMethod) (models.PaymentMethod, error)
	// UpdateById(id int, upPay templates.PaymentMethodRequest) (models.PaymentMethod, error)
	// DeleteById(id int) (gorm.DeletedAt, error)
	// GetAll() ([]templates.PaymentMethodResponse, error)
}
