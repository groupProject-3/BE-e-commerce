package product

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
	return models.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockProdLib struct{}

func (m *MockProdLib) Create(user_id uint, newPro models.Product) (models.Product, error) {
	return models.Product{}, nil
}

func (m *MockProdLib) UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error) {
	return models.Product{}, nil
}

func (m *MockProdLib) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	prod := models.Product{}
	return prod.DeletedAt, nil
}

func (m *MockProdLib) GetAllMe(user_id uint) ([]templates.ProductResponse, error) {
	return []templates.ProductResponse{}, nil
}

