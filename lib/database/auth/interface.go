package auth

import (
	"be/delivery/templates"
	"be/models"
)

type Auth interface {
	Login(UserLogin templates.Userlogin) (models.User, error)
}
