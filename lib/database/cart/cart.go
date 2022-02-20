package cart

import (
	"be/delivery/templates"
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

func (cd *CartDb) GetAll(user_id uint, status string) ([]templates.CartResponse, error) {

	cartRespArr := []templates.CartResponse{}
	log.Info(status)
	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.status = ?", user_id, status).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id, products.qty as Product_qty, carts.qty * products.price as PriceTotal ").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)

	if res.Error != nil || res.RowsAffected == 0 {
		return nil, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}

func (cd *CartDb) GetById(prod_id uint, user_id uint, status string) (templates.CartResponse, error) {
	cartRespArr := templates.CartResponse{}

	res := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status = ?", user_id, prod_id, status).Select("carts.id as ID, carts.created_at as CreatedAt, carts.updated_at as UpdatedAt, carts.qty as Qty, products.price as Price, products.name as Name, products.image as Image, carts.status as Status, carts.product_id as Product_id, products.qty as Product_qty, carts.qty * products.price as PriceTotal").Joins("inner join products on products.id = carts.product_id").Order("products.id asc").Find(&cartRespArr)
	if res.Error != nil /* || res.RowsAffected == 0 */ {
		return cartRespArr, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return cartRespArr, nil
}

func (cd *CartDb) CreateTranx(user_id uint, newCart models.Cart) (models.Cart, error) {
	log.Info(newCart)
	prod_id := newCart.Product_id
	cartInit := models.Cart{}
	row := cd.db.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ? AND carts.status != ?", user_id, prod_id, "payed").Find(&cartInit)

	if row.Error != nil {
		return models.Cart{}, row.Error
	}

	if row.RowsAffected != 0 {
		log.Info(row.RowsAffected)
		res1, err1 := cd.UpdateTranx(prod_id, user_id, templates.CartRequest{Qty: newCart.Qty})
		if err1 != nil {
			return models.Cart{}, err1
		}
		return res1, nil
	}
	log.Info(newCart)
	tx := cd.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	log.Info(newCart)
	if res := tx.Create(&newCart); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Find(&cartInit); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	resProd1 := models.Product{}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	plus := resProd1.Qty - int(cartInit.Qty)
	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Update("qty", plus); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil || resProd1.Qty < 0 {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	return cartInit, tx.Commit().Error
}

func (cd *CartDb) CreateNew(user_id uint, newCart models.Cart) (templates.CartResponse, error) {
	var res1 models.Cart
	var err1 error
	newCart.User_id = user_id
	if res1, err1 = cd.CreateTranx(user_id, newCart); err1 != nil {
		return templates.CartResponse{}, errors.New(gorm.ErrInvalidTransaction.Error())
	}

	res3, err3 := cd.GetById(res1.Product_id, user_id, res1.Status)
	if err3 != nil || res3.Qty != newCart.Qty {
		return templates.CartResponse{}, errors.New(gorm.ErrInvalidTransaction.Error())
	}
	// log.Info(newCart.Qty, res1.Qty, res3.Qty)

	return res3, nil
}

func (cd *CartDb) DeleteTranx(prod_id uint, user_id uint) (models.Cart, error) {

	tx := cd.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return models.Cart{}, err
	}

	cartInit1 := models.Cart{}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Find(&cartInit1); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Delete(&cartInit1); res.RowsAffected == 0 {
		tx.Rollback()
		return models.Cart{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	resProd1 := models.Product{}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	plus := resProd1.Qty + int(cartInit1.Qty)
	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Update("qty", plus); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil || resProd1.Qty < 0 {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	return cartInit1, tx.Commit().Error
}

func (cd *CartDb) DeleteNew(prod_id uint, user_id uint) (models.Cart, error) {
	var res1 models.Cart
	var err1 error

	if res1, err1 = cd.DeleteTranx(prod_id, user_id); err1 != nil {
		return models.Cart{}, errors.New(gorm.ErrInvalidTransaction.Error())
	}

	return res1, nil

}

func (cd *CartDb) UpdateTranx(prod_id uint, user_id uint, upCart templates.CartRequest) (models.Cart, error) {

	tx := cd.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return models.Cart{}, err
	}
	log.Info(upCart.Qty)
	resCart1 := models.Cart{}

	if err := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ? AND carts.status != ?", user_id, prod_id, "payed").Find(&resCart1).Error; err != nil {
		tx.Rollback()
		return models.Cart{}, err
	}
	log.Info(resCart1.Qty)

	cartInit1 := models.Cart{}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Find(&cartInit1); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Delete(&cartInit1); res.RowsAffected == 0 {
		tx.Rollback()
		return models.Cart{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	if res := tx.Create(&models.Cart{User_id: user_id, Product_id: prod_id, Qty: cartInit1.Qty, Status: cartInit1.Status}); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	cartInit2 := models.Cart{}

	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Find(&cartInit2); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	log.Info(cartInit2)
	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, "payed").Updates(models.Cart{Qty: upCart.Qty, Status: upCart.Status}); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	log.Info(cartInit2)
	if res := tx.Model(&models.Cart{}).Where("carts.user_id = ? AND carts.product_id = ?  AND carts.status != ?", user_id, prod_id, cartInit2.Status).Find(&cartInit2); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	log.Info(cartInit2)
	resProd1 := models.Product{}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Find(&resProd1); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}
	var plus int
	log.Info(resProd1.Qty, resCart1.Qty, upCart.Qty)
	if resCart1.Qty < upCart.Qty {
		plus = (resProd1.Qty + int(resCart1.Qty)) - int(upCart.Qty)
	} else {
		plus = (resProd1.Qty + int(resCart1.Qty)) - int(upCart.Qty)
	}
	log.Info(plus)
	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Update("qty", plus); res.Error != nil {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	if res := tx.Model(&models.Product{}).Where("products.id = ?", prod_id).Last(&resProd1); res.Error != nil || resProd1.Qty < 0 {
		tx.Rollback()
		return models.Cart{}, res.Error
	}

	return cartInit2, tx.Commit().Error
}

func (cd *CartDb) UpdateCart(prod_id uint, user_id uint, upCart templates.CartRequest) (templates.CartResponse, error) {
	var res1 models.Cart
	var err1 error

	// if upCart.Status == "" {
	if res1, err1 = cd.UpdateTranx(prod_id, user_id, upCart); err1 != nil {
		return templates.CartResponse{}, err1
	}
	// } else {

	// }

	res3, err3 := cd.GetById(prod_id, user_id, res1.Status)
	if err3 != nil || res3.Qty != upCart.Qty {
		return res3, errors.New(gorm.ErrInvalidTransaction.Error())
	}
	return res3, nil
}
