package auth

import (
	"be/delivery/templates"
	"be/models"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(UserLogin templates.Userlogin) (models.User, error) {
	user := models.User{}
	if err := ad.db.Model(&models.User{}).Where("email = ? AND password = ?", UserLogin.Email, UserLogin.Password).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
