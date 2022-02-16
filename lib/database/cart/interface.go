package cart

import "be/models"

type Cart interface {
	Create(user_id uint, newCart models.Cart) (models.Cart, error)
	// DeleteById(id uint, user_id uint) (gorm.DeletedAt, error)
	// UpdateById(id uint, user_id uint, upCart cart.CartRequest) (models.Cart, error)
	// GetAll(user_id uint) ([]cart.CartResponse, error)
}
