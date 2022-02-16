package order

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type OrderDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OrderDb {
	return &OrderDb{
		db: db,
	}
}

func (od *OrderDb) Create(user_id uint, newOrder models.Order) (models.Order, error) {
	newOrder.User_id = user_id

	if err := od.db.Create(&newOrder).Error; err != nil {
		return models.Order{}, err
	}
	return newOrder, nil
}

func (od *OrderDb) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	order := models.Order{}

	res := od.db.Model(&models.Order{}).Where("id = ? AND user_id = ?", id, user_id).Delete(&order)

	if res.RowsAffected == 0 {
		return order.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return order.DeletedAt, nil
}

func (od *OrderDb) GetById(id uint, user_id uint) (templates.OrderResponse, error) {
	orderResp := templates.OrderResponse{}

	res := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ?", id, user_id).First(&orderResp)

	if res.RowsAffected == 0 {
		return templates.OrderResponse{}, res.Error
	}

	orderDetails := []templates.OrderDetailResponse{}

	resorderDetails := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ?", id, user_id).Select("order_details.id as ID, order_details.created_at as CreatedAt, order_details.updated_at as UpdatedAt, order_details.cart_id as Cart_id, order_details.qty as Qty, order_details.price as Price").Joins("inner join order_details on order_details.order_id = orders.id").Find(&orderDetails)

	if resorderDetails.Error != nil {
		return templates.OrderResponse{}, resorderDetails.Error
	}

	orderResp.OrderDetails = orderDetails

	return orderResp, nil
}
