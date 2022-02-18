package cart

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type Cart interface {
	Create(user_id uint, newCart models.Cart) (models.Cart, error)
	DeleteById(id uint, user_id uint) (gorm.DeletedAt, error)
	UpdateById(id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error)
	GetAll(user_id uint) ([]templates.CartResponse, error)
	FindId(user_id uint, product_id uint) (uint, error)
	GetById(id uint, user_id uint) (templates.CartResponse, error)
}
