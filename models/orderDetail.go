package models

import (
	"be/delivery/controllers/orderDetail"

	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model
	Cart_id  uint
	Order_id uint
	Qty      uint
	Price    uint
}

func (o *OrderDetail) ToOrderDetailResponse() orderDetail.OrderDetailResponse {
	return orderDetail.OrderDetailResponse{
		ID: o.ID,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,

		Cart_id:  o.Cart_id,
		Order_id: o.Order_id,
		Qty:      o.Qty,
		Price:    o.Price,
	}
}
