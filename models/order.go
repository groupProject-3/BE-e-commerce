package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id           uint
	Payment_method_id uint
	Status            string        `gorm:"type:enum('waiting', 'payed','cancel', 'problem');default:'waiting'"`
	OrderDetails      []OrderDetail `gorm:"foreignKey:Order_id"`
}
