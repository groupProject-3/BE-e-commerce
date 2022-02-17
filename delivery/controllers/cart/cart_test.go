package cart

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin templates.Userlogin) (models.User, error) {
	return models.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockCartLib struct{}

func (m *MockCartLib) Create(user_id uint, newCart models.Cart) (models.Cart, error) {
	return models.Cart{}, nil
}

func (m *MockCartLib) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	cart := models.Cart{}

	return cart.DeletedAt, nil
}

func (m *MockCartLib) UpdateById(id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error) {
	return models.Cart{}, nil
}

func (m *MockCartLib) GetAll(user_id uint) ([]templates.CartResponse, error) {
	return []templates.CartResponse{}, nil
}

func (m *MockCartLib) CheckProduct(user_id uint, product_id uint) (bool, error) {
	return false, nil
}

type MockFailCartLib struct{}

func (m *MockFailCartLib) Create(user_id uint, newCart models.Cart) (models.Cart, error) {
	return models.Cart{}, errors.New("")
}

func (m *MockFailCartLib) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	cart := models.Cart{}

	return cart.DeletedAt, errors.New("")
}

func (m *MockFailCartLib) UpdateById(id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error) {
	return models.Cart{}, errors.New("")
}

func (m *MockFailCartLib) GetAll(user_id uint) ([]templates.CartResponse, error) {
	return []templates.CartResponse{}, errors.New("")
}

func (m *MockFailCartLib) CheckProduct(user_id uint, product_id uint) (bool, error) {
	return false, errors.New("")
}
