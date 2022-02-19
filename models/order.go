package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id           uint   `gorm:"primaryKey"`
	Payment_method_id uint   `gorm:"primaryKey"`
	Status            string `gorm:"type:enum('waiting', 'payed','cancel', 'problem');default:'waiting'"`
	PhoneNumber       uint
	OrderDetails      []OrderDetail `gorm:"foreignKey:Order_id"`
}
