package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id       uint `gorm:"primaryKey"`
	Product_id    uint `gorm:"primaryKey"`
	Qty           uint
	Status        string        `gorm:"type:enum('cart', 'order', 'payed','cancel');default:'cart'"`
	Order_details []OrderDetail `gorm:"foreignKey:Cart_id"`
}
