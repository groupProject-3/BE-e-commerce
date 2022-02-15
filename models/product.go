package models

import (
	"be/delivery/controllers/product"

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

func (p *Product) ToProductResponse() product.ProductResponse {
	return product.ProductResponse{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,

		Name:            p.Name,
		Product_type_id: p.Product_type_id,
		Price:           p.Price,
		Qty:             p.Price,
		Description:     p.Description,
	}
}
