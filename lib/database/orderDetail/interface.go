package orderDetail

import "be/models"

type OrderDetail interface {
	Create(newOrderDetail models.OrderDetail) (models.OrderDetail, error)
	// DeleteById(cart_id uint, order_id uint) (gorm.DeletedAt, error)
	// GetAll(order_id uint) ([]orderDetail.OrderDetailResponse, error)
}
