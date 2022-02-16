package prodType

import "be/models"

type ProductType interface {
	Create(proType models.ProductType) (models.ProductType, error)
	//  UpdateById(id int, upPro templates.ProductTypeRequest) (models.ProductType, error)
	//  DeleteById(id int) (gorm.DeletedAt, error)
	//  GetAll() ([]templates.ProductTypeResponse, error)
}
