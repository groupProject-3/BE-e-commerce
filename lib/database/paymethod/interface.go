package paymethod

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type PaymentMethod interface {
	Create(pm models.PaymentMethod) (models.PaymentMethod, error)
	UpdateById(id int, upPay templates.PaymentMethodRequest) (models.PaymentMethod, error)
	DeleteById(id int) (gorm.DeletedAt, error)
	GetAll() ([]templates.PaymentMethodResponse, error)
}
