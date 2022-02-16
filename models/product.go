package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	User_id         uint
	Product_type_id uint
	Name            string `gorm:"not null;type:varchar(100)"`
	Price           int    `gorm:"not null"`
	Qty             int    `gorm:"not null"`
	Description     string
	Carts           []Cart `gorm:"foreignKey:Product_id"`
}
