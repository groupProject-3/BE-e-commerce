package cart

import (
	"be/delivery/templates"
	"be/lib/database/product"
	"be/models"
	"errors"

	"github.com/labstack/gommon/log"
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

func (cd *CartDb) DeleteById(pro_id uint, user_id uint) (gorm.DeletedAt, error) {

	cart := models.Cart{}

	res := cd.db.Model(&models.Cart{}).Where("product_id = ? AND user_id = ?", pro_id, user_id).Delete(&cart)

	if res.RowsAffected == 0 {
		return cart.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cart.DeletedAt, nil
}

func (cd *CartDb) UpdateById(prod_id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error) {

	cartInit := models.Cart{}

	if res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Find(&cartInit); res.Error != nil {
		return models.Cart{}, res.Error
	}

	if _, err := cd.DeleteById(prod_id, user_id); err != nil {
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

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ?", user_id).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id, products.qty as Product_qty, carts.qty * products.price as PriceTotal ").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)

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

func (cd *CartDb) GetById(prod_id uint, user_id uint) (templates.CartResponse, error) {
	cartRespArr := templates.CartResponse{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id, products.qty as Product_qty").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)
	if res.Error != nil /* || res.RowsAffected == 0 */ {
		return cartRespArr, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}

func (cd *CartDb) CreateNew(user_id uint, newCart models.Cart) (templates.CartResponse, error) {
	prod_id := newCart.Product_id
	cartinit := models.Cart{}
	row := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Find(&cartinit)

	if row.Error != nil {
		return templates.CartResponse{}, row.Error
	}
	if row.RowsAffected != 0 {
		res1, err1 := cd.GetById(prod_id, user_id)
		if err1 != nil {
			return templates.CartResponse{}, err1
		}
		log.Info(res1.Qty)
		if _, err := cd.UpdateById(prod_id, user_id, templates.CartRequest{Qty: newCart.Qty}); err != nil {
			return templates.CartResponse{}, err
		}

		res2, err2 := product.New(cd.db).GetById(int(prod_id))

		if err2 != nil {
			return templates.CartResponse{}, err2
		}

		log.Info(res2.Qty, newCart.Qty)
		if _, err := product.New(cd.db).UpdateByIdAll(int(prod_id), templates.ProductRequest{Qty: (res2.Qty + ((res1.Qty) - (newCart.Qty)))}); err != nil {
			return templates.CartResponse{}, err
		}

		res3, err3 := cd.GetById(prod_id, user_id)
		if err3 != nil {
			return templates.CartResponse{}, err3
		}
		log.Info(res3.Qty)
		return res3, nil
	}

	if _, err := cd.Create(user_id, newCart); err != nil {
		return templates.CartResponse{}, err
	}

	res2, err2 := product.New(cd.db).GetById(int(prod_id))

	if err2 != nil {
		return templates.CartResponse{}, err2
	}

	if _, err := product.New(cd.db).UpdateByIdAll(int(prod_id), templates.ProductRequest{Qty: (res2.Qty - (newCart.Qty))}); err != nil {
		return templates.CartResponse{}, err
	}

	res3, err3 := cd.GetById(prod_id, user_id)
	if err3 != nil {
		return templates.CartResponse{}, err3
	}

	res3.PriceTotal = res3.Qty * uint(res3.Price)

	return res3, nil
}

func (cd *CartDb) DeleteNew(prod_id uint, user_id uint) (gorm.DeletedAt, error) {

	cart := models.Cart{}

	res1, err1 := cd.GetById(prod_id, user_id)
	if err1 != nil {
		return cart.DeletedAt, err1
	}
	log.Info(product.New(cd.db).GetById(int(prod_id)))
	res := cd.db.Model(&models.Cart{}).Where("product_id = ? AND user_id = ?", prod_id, user_id).Delete(&cart)

	if res.RowsAffected == 0 {
		return cart.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}
	log.Info(product.New(cd.db).GetById(int(prod_id)))
	res2, err2 := product.New(cd.db).GetById(int(prod_id))
	log.Info(res2)
	if err2 != nil {
		return cart.DeletedAt, err2
	}
	log.Info(product.New(cd.db).GetById(int(prod_id)))
	if _, err := product.New(cd.db).UpdateByIdAll(int(prod_id), templates.ProductRequest{Qty: (res2.Qty + (res1.Qty))}); err != nil {
		return cart.DeletedAt, err
	}
	log.Info(product.New(cd.db).GetById(int(prod_id)))
	return cart.DeletedAt, nil
}

func (cd *CartDb) UpdateNew(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error) {

	res1, err1 := cd.GetById(prod_id, user_id)
	if err1 != nil {
		return templates.CartResponse{}, err1
	}
	log.Info(res1.Qty, res1.Product_qty)
	if _, err := cd.UpdateById(prod_id, user_id, templates.CartRequest{Qty: upCart.Qty, Status: "order"}); err != nil {
		return templates.CartResponse{}, err
	}

	res2, err2 := product.New(cd.db).GetById(int(prod_id))

	if err2 != nil {
		return templates.CartResponse{}, err2
	}
	if _, err := product.New(cd.db).UpdateByIdAll(int(prod_id), templates.ProductRequest{Qty: (res2.Qty + ((res1.Qty) - (upCart.Qty)))}); err != nil {
		return templates.CartResponse{}, err
	}

	res3, err3 := cd.GetById(prod_id, user_id)
	if err3 != nil {
		return templates.CartResponse{}, err3
	}

	res3.PriceTotal = res3.Qty * uint(res3.Price)

	return res3, nil
}

func (cd *CartDb) UpdateTranx(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error) {

	tx := cd.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return templates.CartResponse{}, err
	}
	log.Info(upCart.Qty)
	resCart1 := models.Cart{}

	if err := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Find(&resCart1).Error; err != nil {
		tx.Rollback()
		return templates.CartResponse{}, err
	}
	log.Info(resCart1.Qty)

	cartInit1 := models.Cart{}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Find(&cartInit1); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Delete(&cartInit1); res.RowsAffected == 0 {
		tx.Rollback()
		return templates.CartResponse{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	if res := tx.Create(&models.Cart{User_id: user_id, Product_id: prod_id, Qty: cartInit1.Qty, Status: cartInit1.Status}); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	cartInit2 := models.Cart{}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Find(&cartInit2); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?", user_id, prod_id).Updates(models.Cart{Qty: upCart.Qty, Status: upCart.Status}).First(&cartInit2); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	resProd1 := models.Product{}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Updates(models.Product{Qty: resProd1.Qty + int(resCart1.Qty-upCart.Qty)}); res.Error != nil {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil || resProd1.Qty < 0 {
		tx.Rollback()
		return templates.CartResponse{}, res.Error
	}

	// log.Info(tx.Commit().Error)

	return templates.CartResponse{}, tx.Commit().Error
}

func (cd *CartDb) UpdateCart(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error) {

	if _, err1 := cd.UpdateTranx(prod_id, user_id, upCart); err1 != nil {
		return templates.CartResponse{}, err1
	}

	res3, err3 := cd.GetById(prod_id, user_id)
	if err3 != nil || res3.Qty != upCart.Qty {
		return res3, errors.New(gorm.ErrInvalidTransaction.Error())
	}
	return res3, nil
}
