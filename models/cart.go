package models

import (
	"be/delivery/controllers/cart"

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

func (c *Cart) ToCartResponse() cart.CartResponse {
	return cart.CartResponse{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,

		Product_id: c.Product_id,
		Qty:        c.Qty,
		Status:     c.Status,
	}
}
