package models

import (
	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model

	Name     string
	Products []Product `gorm:"foreignKey:Product_type_id"`
}
