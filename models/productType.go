package models

import (
	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model

	Name     string    `gorm:"unique;index;not null;type:varchar(100)"`
	Products []Product `gorm:"foreignKey:Product_type_id"`
}
