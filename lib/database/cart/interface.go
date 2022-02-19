package cart

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type Cart interface {
	Create(user_id uint, newCart models.Cart) (models.Cart, error)
	DeleteById(pro_id uint, user_id uint) (gorm.DeletedAt, error)
	UpdateById(prod_id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error)
	GetAll(user_id uint) ([]templates.CartResponse, error)
	FindId(user_id uint, product_id uint) (uint, error)
	GetById(prod_id uint, user_id uint) (templates.CartResponse, error)
	CreateNew(user_id uint, newCart models.Cart) (templates.CartResponse, error)
	DeleteNew(prod_id uint, user_id uint) (gorm.DeletedAt, error)
	UpdateCart(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error)
}
