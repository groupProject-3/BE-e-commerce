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

	cartInit := models.Cart{}

	if res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.id = ?", user_id, id).Find(&cartInit); res.Error != nil {
		return models.Cart{}, res.Error
	}

	if _, err := cd.DeleteById(id, user_id); err != nil {
		return models.Cart{}, err
	}
	// log.Info(cd.GetAll(1))
	cd.db.Create(&models.Cart{User_id: user_id, Product_id: cartInit.Product_id, Qty: cartInit.Qty, Status: cartInit.Status})
	// log.Info(cd.GetAll(1))
	cartInit2 := models.Cart{}
	cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, cartInit.Product_id).Find(&cartInit2)
	// log.Info(cartInit2)
	// log.Info(cd.GetAll(1))
	err := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.id = ?", user_id, cartInit2.ID).Updates(models.Cart{Qty: upCart.Qty, Status: upCart.Status}).First(&cartInit2)

	if err.Error != nil {
		return models.Cart{}, err.Error
	}
	return cartInit, nil
}

func (cd *CartDb) GetAll(user_id uint) ([]templates.CartResponse, error) {
	cartRespArr := []templates.CartResponse{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ?", user_id).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)

	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}

func (cd *CartDb) FindId(user_id uint, product_id uint) (uint, error) {

	cart := models.Cart{}

	cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, product_id).Find(&cart)

	if cart.ID == 0 {
		return 0, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cart.ID, nil
}

func (cd *CartDb) GetById(id uint, user_id uint) (templates.CartResponse, error) {
	cartRespArr := templates.CartResponse{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.id = ?", user_id, id).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)
	if res.Error != nil /* || res.RowsAffected == 0 */ {
		return cartRespArr, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}
