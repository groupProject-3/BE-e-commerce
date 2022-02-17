package product

import (
	"be/delivery/templates"
	"be/models"
	"errors"

	"gorm.io/gorm"
)

type ProductDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (pd *ProductDb) Create(user_id uint, newPro models.Product) (models.Product, error) {
	newPro.User_id = user_id
	if err := pd.db.Create(&newPro).Error; err != nil {
		return newPro, err
	}
	return newPro, nil
}

func (pd *ProductDb) UpdateById(id int, user_id uint, upPro templates.ProductRequest) (models.Product, error) {
	pro := models.Product{}

	res := pd.db.Model(&models.Product{}).Where("id = ? AND user_id = ?", id, user_id).Updates(models.Product{Name: upPro.Name, Image: pro.Image, Product_type_id: upPro.Product_type_id, Qty: upPro.Qty, Price: upPro.Price, Description: upPro.Description}).First(&pro)

	if res.RowsAffected == 0 {
		return models.Product{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro, nil
}

func (pd *ProductDb) DeleteById(id int, user_id uint) (gorm.DeletedAt, error) {
	pro := models.Product{}

	res := pd.db.Model(&models.Product{}).Where("id = ? AND user_id = ?", id, user_id).Delete(&pro)

	if res.RowsAffected == 0 {
		return pro.DeletedAt, errors.New(gorm.ErrRecordNotFound.Error())
	}

	return pro.DeletedAt, nil
}

func (pd *ProductDb) GetAllMe(user_id uint) ([]templates.ProductResponse, error) {
	pro := []templates.ProductResponse{}

	res := pd.db.Model(&models.Product{}).Where("products.user_id = ?", user_id).Select("products.id as ID, products.created_at as CreatedAt, products.updated_at as UpdatedAt, products.name as Name, products.image as Image, products.price as Price, products.qty as Qty, products.description as Description, product_types.name as Product_type_name").Joins("inner join product_types on product_types.id = products.product_type_id").Order("products.id asc").Find(&pro)

	if res.RowsAffected == 0 {
		return pro, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return pro, nil
}

func (pd *ProductDb) GetByIdMe(id int, user_id uint) (templates.ProductResponse, error) {
	pro := templates.ProductResponse{}

	res := pd.db.Model(&models.Product{}).Where("products.id = ? AND products.user_id = ?", id, user_id).Select("products.id as ID, products.created_at as CreatedAt, products.updated_at as UpdatedAt, products.name as Name, products.image as Image, products.price as Price, products.qty as Qty, products.description as Description, product_types.name as Product_type_name").Joins("inner join product_types on product_types.id = products.product_type_id").Order("products.id asc").First(&pro)

	if res.Error != nil {
		return pro, res.Error
	}
	return pro, nil
}

func (pd *ProductDb) GetAll() ([]templates.ProductResponse, error) {
	pro := []templates.ProductResponse{}

	res := pd.db.Model(&models.Product{}).Select("products.id as ID, products.created_at as CreatedAt, products.updated_at as UpdatedAt, products.name as Name, products.image as Image, products.price as Price, products.qty as Qty, products.description as Description, product_types.name as Product_type_name").Joins("inner join product_types on product_types.id = products.product_type_id").Order("products.id asc").Find(&pro)

	if res.RowsAffected == 0 {
		return pro, errors.New(gorm.ErrRecordNotFound.Error())
	}
	return pro, nil
}

func (pd *ProductDb) GetById(id int) (templates.ProductResponse, error) {
	pro := templates.ProductResponse{}

	res := pd.db.Model(&models.Product{}).Where("products.id = ?", id).Select("products.id as ID, products.created_at as CreatedAt, products.updated_at as UpdatedAt, products.name as Name, products.image as Image, products.price as Price, products.qty as Qty, products.description as Description, product_types.name as Product_type_name").Joins("inner join product_types on product_types.id = products.product_type_id").Order("products.id asc").First(&pro)

	if res.Error != nil {
		return pro, res.Error
	}
	return pro, nil
}
