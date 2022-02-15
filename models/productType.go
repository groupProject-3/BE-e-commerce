package models

import (
	producttype "be/delivery/controllers/productType"

	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model

	Name string
	Products []Product `gorm:"foreignKey:Product_type_id"` 
}

func (p *ProductType) ToProductTypeResponse() producttype.ProductTypeResponse{
	return producttype.ProductTypeResponse{
		ID: p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		
		Name: p.Name,
	}
}