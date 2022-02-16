package order

import "be/models"

type Order interface {
	Create(user_id uint, newOrder models.Order) (models.Order, error)
	// DeleteById(id uint, user_id uint) (gorm.DeletedAt, error)
	// GetById(id uint, user_id uint) (order.OrderResponse, error)
}
