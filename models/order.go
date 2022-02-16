package models

import (
	"be/delivery/controllers/order"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User_id           uint
	Payment_method_id uint
	Status            string        `gorm:"type:enum('waiting', 'payed','cancel', 'problem');default:'waiting'"`
	OrderDetails      []OrderDetail `gorm:"foreignKey:Order_id"`
}

func (o *Order) ToOrderResponse() order.OrderResponse {
	return order.OrderResponse{
		ID:        o.ID,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,

		Payment_method_id: o.Payment_method_id,
		Status:            o.Status,
	}
}
