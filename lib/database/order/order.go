package order

import (
	// "be/lib/database/orderDetail"
	"be/delivery/controllers/order"
	"be/delivery/controllers/orderDetail"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type OrderDb struct {
	db *gorm.DB
	// orderDetai *orderDetail
}

func New(db *gorm.DB) *OrderDb {
	return &OrderDb{
		db: db,
	}
}

func (od *OrderDb) Create(user_id uint, newOrder models.Order) (models.Order, error) {
	newOrder.User_id = user_id

	sumQty := od.db.Model(&models.OrderDetail{}).Where("order_id = ?", newOrder.ID).Select("sum(qty)").Row().Scan(&newOrder.Total_qty)
	if sumQty.Error() != "" {
		return models.Order{}, errors.New(sumQty.Error())
	}
	sumPrice := od.db.Model(&models.OrderDetail{}).Where("order_id = ?", newOrder.ID).Select("sum(qty)").Row().Scan(&newOrder.Total_price)
	if sumPrice.Error() != "" {
		return models.Order{}, errors.New(sumPrice.Error())
	}
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

func (od *OrderDb) GetById(id uint, user_id uint) (order.OrderResponse, error) {
	orderResp := order.OrderResponse{}

	res := od.db.Model(&models.Order{}).Where("id = ? AND user_id = ?", id, user_id).First(&orderResp)

	if res.RowsAffected == 0 {
		return order.OrderResponse{}, res.Error
	}

	orderDetails := []orderDetail.OrderDetailResponse{}

	resorderDetails := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ?", id, user_id).Select("order_details.id as ID, order_details.created_at as CreatedAt, order_details.updated_at as UpdatedAt, order_details.cart_id as Cart_id, order_details.qty as Qty, order_details.price as Price").Joins("innter join order_details on order_details.order_id = orders.id").Find(&orderDetails)

	if resorderDetails.Error != nil {
		return order.OrderResponse{}, resorderDetails.Error
	}

	orderResp.OrderDetails = orderDetails

	return orderResp, nil
}
