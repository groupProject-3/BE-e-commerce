package orderDetail

import (
	"be/delivery/controllers/orderDetail"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type OrderDetailDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OrderDetailDb {
	return &OrderDetailDb{
		db: db,
	}
}

func (od *OrderDetailDb) Create(newOrderDetail models.OrderDetail) (models.OrderDetail, error) {
	if err := od.db.Create(&newOrderDetail).Error; err != nil {
		return models.OrderDetail{}, err
	}

	return newOrderDetail, nil
}

func (od *OrderDetailDb) DeleteById(id uint) (gorm.DeletedAt, error) {
	orderDetail := models.OrderDetail{}

	res := od.db.Model(&models.OrderDetail{}).Where("id = ?", id).Delete(&orderDetail)

	if res.RowsAffected == 0 {
		return orderDetail.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return orderDetail.DeletedAt, nil
}

func (od *OrderDetailDb) GetAll(order_id uint) ([]orderDetail.OrderDetailResponse, error) {
	orderDetailRespArr := []orderDetail.OrderDetailResponse{}

	res := od.db.Model(&models.OrderDetail{}).Where("order_id = ?", order_id).Find(&orderDetailRespArr)

	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return orderDetailRespArr, nil
}
