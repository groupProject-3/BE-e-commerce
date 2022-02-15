package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User_id       uint
	Product_id    uint
	Qty           uint
	Price         uint
	Status        bool
	Order_details []OrderDetail `gorm:"foreignKey:Cart_id"`
}
