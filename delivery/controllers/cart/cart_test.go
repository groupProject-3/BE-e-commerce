package cart

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
	return models.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockCartLib struct{}

func 