package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	User_id       uint
	Product_id    uint
	Qty           uint
	Status        string        `gorm:"type:enum('cart', 'order', 'payed','cancel');default:'cart'"`
	Order_details []OrderDetail `gorm:"foreignKey:Cart_id"`
}
