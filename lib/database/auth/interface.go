package auth

import (
	"be/delivery/controllers/auth"
	"be/models"
)

type Auth interface {
	Login(UserLogin auth.Userlogin) (models.User, error)
}
