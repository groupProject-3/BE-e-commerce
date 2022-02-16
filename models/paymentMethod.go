package models

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model

	Name   string
	Orders []Order `gorm:"foreignKey:Payment_method_id"`
}
