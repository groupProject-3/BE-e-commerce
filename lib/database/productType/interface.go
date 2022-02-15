package producttype

import "be/models"

type ProductType interface {
	Create(proType models.ProductType) (models.ProductType, error)
	//  UpdateById(id int, upPro producttype.ProductTypeRequest) (models.ProductType, error)
	//  DeleteById(id int) (gorm.DeletedAt, error)
	//  GetAll() ([]producttype.ProductTypeResponse, error)
}
