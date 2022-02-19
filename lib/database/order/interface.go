package order

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type Order interface {
	Create(user_id uint, newOrder models.Order) (models.Order, error)
	DeleteById(user_id uint) (gorm.DeletedAt, error)
	GetById(id uint, user_id uint) (templates.OrderResponse, error)
	Update(user_id int, upOr templates.OrderRequest) (models.Order, error)
}
