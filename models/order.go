package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User_id           uint
	Payment_method_id uint
	Total_qty         uint
	Total_price       uint
	Status            string        `gorm:"type:enum('waiting', 'payed','cancel', 'problem');default:'waiting'"`
	Order_details     []OrderDetail `gorm:"foreignKey:Order_id"`
}
