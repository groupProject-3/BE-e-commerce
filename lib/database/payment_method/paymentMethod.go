package paymentmethod

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type PaymentMethodDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PaymentMethodDb {
	return &PaymentMethodDb{
		db: db,
	}
}

func (pd *PaymentMethodDb) Create(pm models.PaymentMethod) (models.PaymentMethod, error) {
	if err := pd.db.Create(&pm).Error; err != nil {
		return models.PaymentMethod{}, err
	}
	return pm, nil
}

func (pd *PaymentMethodDb) UpdateById(id int, upPay templates.PaymentMethodRequest) (models.PaymentMethod, error) {
	pm := models.PaymentMethod{}

	res := pd.db.Model(&models.PaymentMethod{}).Where("id = ?", id).Updates(models.PaymentMethod{Name: upPay.Name}).First(&pm)

	if res.RowsAffected == 0 {
		return models.PaymentMethod{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pm, nil
}

func (pd *PaymentMethodDb) DeleteById(id int) (gorm.DeletedAt, error) {
	pm := models.PaymentMethod{}

	res := pd.db.Model(&models.PaymentMethod{}).Where("id = ?", id).Delete(&pm)

	if res.RowsAffected == 0 {
		return pm.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pm.DeletedAt, nil
}

func (pd *PaymentMethodDb) GetAll() ([]templates.PaymentMethodResponse, error) {
	pmRespArr := []templates.PaymentMethodResponse{}

	res := pd.db.Model(&models.PaymentMethod{}).Find(&pmRespArr)

	if res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pmRespArr, nil
}
