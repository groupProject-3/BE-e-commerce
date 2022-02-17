package product

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type Product interface {
	Create(user_id uint, newPro models.Product) (models.Product, error)
	UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error)
	DeleteById(id int, user_id uint) (gorm.DeletedAt, error)
	GetAllMe(user_id uint) ([]templates.ProductResponse, error)
	GetByIdMe(id int, user_id uint) (templates.ProductResponse, error)
	GetAll() ([]templates.ProductResponse, error)
	GetById(id int) (templates.ProductResponse, error)
}
