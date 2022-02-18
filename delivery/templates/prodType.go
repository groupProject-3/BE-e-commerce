package templates

import (
	"be/models"
	"time"
)

type ProductTypeRequest struct {
	Name string `json:"name"`
}

type ProductTypeResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`
}

type GetProdTypeResponseFormat struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    models.ProductType `json:"data"`
}
