package product

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type ProductDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (pd *ProductDb) Create(user_id uint, newPro models.Product) (models.Product, error) {
	newPro.User_id = user_id
	if err := pd.db.Create(&newPro).Error; err != nil {
		return newPro, err
	}
	return newPro, nil
}

func (pd *ProductDb) UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error) {
	pro := models.Product{}

	res := pd.db.Model(&models.Product{}).Where("id = ? AND user_id = ?", id, user_id).Updates(models.Product{Name: upPro.Name, Product_type_id: upPro.Product_type_id, Qty: upPro.Qty, Price: upPro.Price, Description: upPro.Description}).First(&pro)

	if res.RowsAffected == 0 {
		return models.Product{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro, nil
}

func (pd *ProductDb) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	pro := models.Product{}

	res := pd.db.Model(&models.Product{}).Where("id = ? AND user_id = ?", id, user_id).Delete(&pro)

	if res.RowsAffected == 0 {
		return pro.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro.DeletedAt, nil
}

func (pd *ProductDb) GetAllMe(user_id uint) ([]models.Product, error) {
	pro := []models.Product{}

	res := pd.db.Model(&models.Product{}).Where("user_id = ?", user_id).Find(&pro)

	if res.RowsAffected == 0 {
		return pro, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return pro, nil
}

func (pd *ProductDb) GetByIdMe(id int, user_id uint) (models.Product, error) {
	pro := models.Product{}

	if err := pd.db.Model(&models.Product{}).Where("id = ? AND user_id = ?", id, user_id).First(&pro).Error; err != nil {
		return pro, err
	}
	return pro, nil
}

func (pd *ProductDb) GetAll() ([]models.Product, error) {
	pro := []models.Product{}

	res := pd.db.Model(&models.Product{}).Find(&pro)

	if res.RowsAffected == 0 {
		return pro, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return pro, nil
}

func (pd *ProductDb) GetById(id int) (models.Product, error) {
	pro := models.Product{}

	if err := pd.db.Model(&models.Product{}).Where("id = ?", id).First(&pro).Error; err != nil {
		return pro, err
	}
	return pro, nil
}
