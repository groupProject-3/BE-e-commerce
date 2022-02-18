package models

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model

	Name   string  `gorm:"unique;index;not null;type:varchar(100)"`
	Orders []Order `gorm:"foreignKey:Payment_method_id"`
}
