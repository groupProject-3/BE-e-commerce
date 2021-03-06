package order

import (
	"be/delivery/templates"
	"be/lib/database/cart"
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

func (od *OrderDb) GetById(user_id uint) (templates.OrderResponse, error) {
	orderResp := templates.OrderResponse{}

	res := od.db.Model(&models.Order{}).Where("orders.user_id = ?", user_id).Select("orders.id as ID, orders.created_at as CreatedAt, orders.updated_at as UpdatedAt, users.name as Name, orders.payment_method_id as Payment_method_id, orders.status as Status").Joins("inner join users on users.id = orders.user_id").Last(&orderResp)

	if res.RowsAffected == 0 {
		return templates.OrderResponse{}, res.Error
	}

	orderDetails := []templates.CartResponse{}

	resorderDetails := od.db.Model(&models.Order{}).Where("orders.id = ? AND orders.user_id = ? AND carts.status = ? AND carts.deleted_at IS NULL", orderResp.ID, user_id, "order").Select("orders.id as ID, orders.created_at as CreatedAt, orders.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id, products.qty as Product_qty, carts.qty * products.price as PriceTotal").Joins("inner join carts on carts.user_id = orders.user_id").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&orderDetails)

	if resorderDetails.Error != nil {
		return templates.OrderResponse{}, resorderDetails.Error
	}

	orderResp.OrderDetails = orderDetails

	return orderResp, nil
}

func (pd *OrderDb) Update(user_id int, upOr templates.OrderRequest) (models.Order, error) {

	pro := models.Order{}

	res := pd.db.Model(&models.Order{}).Where("user_id = ?", user_id).Updates(models.Order{Payment_method_id: upOr.Payment_method_id, Status: upOr.Status, PhoneNumber: upOr.PhoneNumber, Address: upOr.Address}).Find(&pro)

	if res.RowsAffected == 0 {
		return models.Order{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	cartRespArr := []models.Cart{}

	res1 := pd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.status = ?", user_id, "order").Find(&cartRespArr)

	if res1.Error != nil {
		return models.Order{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	for i := 0; i < len(cartRespArr); i++ {
		if _, err := cart.New(pd.db).UpdateTranx(cartRespArr[i].Product_id, uint(user_id), templates.CartRequest{Status: "payed"}); err != nil {
			return models.Order{}, err
		}
	}

	res2 := pd.db.Model(&models.Order{}).Where("user_id = ?", user_id).Delete(&pro)

	if res2.RowsAffected == 0 {
		return models.Order{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro, nil
}
