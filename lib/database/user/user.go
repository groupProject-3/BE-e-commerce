package user

import (
	"be/delivery/templates"
	"be/models"
	"errors"

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
	if err := ud.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ud *UserDb) GetById(id int) (models.User, error) {
	user := models.User{}

	res := ud.db.Model(&models.User{}).Where("id = ?", id).First(&user)

	if res.RowsAffected == 0 {
		return models.User{}, res.Error
	}
	return user, nil
}

func (ud *UserDb) UpdateById(id int, userUp templates.UserRequest) (models.User, error) {

	user := models.User{}

	res := ud.db.Model(&models.User{Model: gorm.Model{ID: uint(id)}}).Updates(models.User{Name: userUp.Name, Email: userUp.Email, Password: userUp.Password}).First(&user)

	if res.RowsAffected == 0 {
		return models.User{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return user, nil
}

func (ud *UserDb) DeleteById(id int) (gorm.DeletedAt, error) {
	user := models.User{}

	res := ud.db.Model(&models.User{}).Where("id = ?", id).Delete(&user)
	if res.RowsAffected == 0 {
		return user.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return user.DeletedAt, nil
}

func (ud *UserDb) GetAll() ([]templates.UserResponse, error) {
	userRespArr := []templates.UserResponse{}

	res := ud.db.Model(&models.User{}).Find(&userRespArr)
	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return userRespArr, nil
}