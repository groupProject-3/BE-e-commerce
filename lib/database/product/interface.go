package product

import "be/models"

type Product interface {
	Create(user_id uint, newPro models.Product) (models.Product, error)
	// UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error)
	// DeleteById(id int, user_id uint) (gorm.DeletedAt, error)
	// GetAllMe(user_id uint) ([]models.Product, error)
	// GetByIdMe(id int, user_id uint) (models.Product, error)
	// GetAll() ([]models.Product, error)
	// GetById(id int) (models.Product, error)
}
