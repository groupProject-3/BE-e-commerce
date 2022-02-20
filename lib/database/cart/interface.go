package cart

import (
	"be/delivery/templates"
	"be/models"
)

type Cart interface {
	GetAll(user_id uint, status string) ([]templates.CartResponse, error)
	CreateNew(user_id uint, newCart models.Cart) (templates.CartResponse, error)
	DeleteNew(prod_id uint, user_id uint) (models.Cart, error)
	UpdateCart(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error)
}
