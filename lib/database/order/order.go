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

func (od *OrderDb) DeleteById(user_id uint) (gorm.DeletedAt, error) {
	order := models.Order{}

	res := od.db.Model(&models.Order{}).Where("user_id = ?", user_id).Delete(&order)

	if res.RowsAffected == 0 {
		return order.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return order.DeletedAt, nil
}

func (od *OrderDb) GetById(id uint, user_id uint) (templates.OrderResponse, error) {
	orderResp := templates.OrderResponse{}

	res := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ?", id, user_id).Select("orders.id as ID, orders.created_at as CreatedAt, orders.updated_at as UpdatedAt, users.name as Name, orders.payment_method_id as Payment_method_id, orders.status as Status").Joins("inner join users on users.id = orders.user_id").First(&orderResp)

	if res.RowsAffected == 0 {
		return templates.OrderResponse{}, res.Error
	}

	orderDetails := []templates.CartResponse{}

	resorderDetails := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ? AND carts.status = ? AND carts.deleted_at IS NULL", id, user_id, "order").Select("orders.id as ID, orders.created_at as CreatedAt, orders.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id").Joins("inner join carts on carts.user_id = orders.user_id").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&orderDetails)

	if resorderDetails.Error != nil {
		return templates.OrderResponse{}, resorderDetails.Error
	}

	orderResp.OrderDetails = orderDetails

	return orderResp, nil
}

func (pd *OrderDb) Update(user_id int, upOr templates.OrderRequest) (models.Order, error) {

	pro := models.Order{}

	res := pd.db.Model(&models.Order{}).Where("user_id = ?", user_id).Updates(models.Order{Payment_method_id: upOr.Payment_method_id, Status: upOr.Status, PhoneNumber: upOr.PhoneNumber}).Find(&pro)

	if res.RowsAffected == 0 {
		return models.Order{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro, nil
}
