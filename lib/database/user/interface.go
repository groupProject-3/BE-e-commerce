package user

import "be/models"

type User interface {
	Create(user models.User) (models.User, error)
	// GetById(id int) (models.User, error)
	// UpdateById(id int, userUp user.UserRequest) (models.User, error)
	// DeleteById(id int) (gorm.DeletedAt, error)
	// GetAll() ([]user.UserResponse, error)
}
