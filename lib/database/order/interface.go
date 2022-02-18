package order

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type Order interface {
	Create(user_id uint, newOrder models.Order) (models.Order, error)
	DeleteById(id uint, user_id uint) (gorm.DeletedAt, error)
	GetById(id uint, user_id uint) (templates.OrderResponse, error)
}
