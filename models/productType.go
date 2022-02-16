package models

import (
	prodType "be/delivery/controllers/prodType"

	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model

	Name string
	Products []Product `gorm:"foreignKey:Product_type_id"` 
}

func (p *ProductType) ToProductTypeResponse() prodType.ProductTypeResponse{
	return prodType.ProductTypeResponse{
		ID: p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		
		Name: p.Name,
	}
}