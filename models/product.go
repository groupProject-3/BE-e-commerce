package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	User_id         uint
	Product_type_id uint
	Name            string `gorm:"not null;type:varchar(100)"`
	Image           string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
	Price           uint    `gorm:"not null"`
	Qty             uint    `gorm:"not null"`
	Description     string
	Carts           []Cart `gorm:"foreignKey:Product_id"`
}
