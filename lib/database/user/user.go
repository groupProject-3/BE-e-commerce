package user

import (
	"be/models"

	"gorm.io/gorm"
)

type UserDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (ud *UserDb) Create(user models.User) (models.User, error) {
	if err := ud.db.Create(&user).Error; err != nil{
		return models.User{}, err
	}
	return user, nil
}

// func (ud *UserDb) 