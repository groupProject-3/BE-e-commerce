package cart

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type CartDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CartDb {
	return &CartDb{
		db: db,
	}
}

func (cd *CartDb) Create(user_id uint, newCart models.Cart) (models.Cart, error) {
	newCart.User_id = user_id
	if err := cd.db.Create(&newCart).Error; err != nil {
		return models.Cart{}, err
	}
	return newCart, nil
}

func (cd *CartDb) DeleteById(id uint, user_id uint) (gorm.DeletedAt, error) {
	cart := models.Cart{}

	res := cd.db.Model(&models.Cart{}).Where("id = ? AND user_id = ?", id, user_id).Delete(&cart)

	if res.RowsAffected == 0 {
		return cart.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cart.DeletedAt, nil
}

func (cd *CartDb) UpdateById(id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error) {

	if _, err := cd.DeleteById(id, user_id); err != nil {
		return models.Cart{}, err
	}
	cart := models.Cart{Product_id: upCart.Product_id, Qty: upCart.Qty, Status: upCart.Status}
	res, err := cd.Create(user_id, cart)

	if err != nil {
		return models.Cart{}, err
	}
	return res, nil
}

func (cd *CartDb) GetAll(user_id uint) ([]templates.CartResponse, error) {
	cartRespArr := []templates.CartResponse{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ?", user_id).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.name as Name, products.image as Image, carts.status as Status").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)

	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}

func (cd *CartDb) CheckProduct(user_id uint, product_id uint) (bool, error) {

	cart := models.Cart{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, product_id)

	if cart.ID == 0 {
		return false, errors.New(gorm.ErrRecordNotFound.Error())
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
