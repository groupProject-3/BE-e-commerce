package prodType

import (
	prodType "be/delivery/controllers/prodType"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type ProductTypeDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductTypeDb {
	return &ProductTypeDb{
		db: db,
	}
}

func (pd *ProductTypeDb) Create(proType models.ProductType) (models.ProductType, error) {
	if err := pd.db.Create(&proType).Error; err != nil {
		return models.ProductType{}, err
	}

	return proType, nil
}

func (pd *ProductTypeDb) UpdateById(id int, upPro prodType.ProductTypeRequest) (models.ProductType, error) {

	pro := models.ProductType{}

	res := pd.db.Model(&models.ProductType{}).Where("id = ?", id).Updates(models.ProductType{Name: upPro.Name}).First(&pro)

	if res.RowsAffected == 0 {
		return models.ProductType{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro, nil
}

func (pd *ProductTypeDb) DeleteById(id int) (gorm.DeletedAt, error) {

	pro := models.ProductType{}

	res := pd.db.Model(&models.ProductType{}).Where("id = ?", id).Delete(&pro)

	if res.RowsAffected == 0 {
		return pro.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro.DeletedAt, nil
}

func (pd *ProductTypeDb) GetAll() ([]prodType.ProductTypeResponse, error) {
	proRespArr := []prodType.ProductTypeResponse{}

	res := pd.db.Model(&models.ProductType{}).Find(&proRespArr)

	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return proRespArr, nil
}
