package user

import "be/models"

type User interface {
	Create(user models.User) (models.User, error)
}
