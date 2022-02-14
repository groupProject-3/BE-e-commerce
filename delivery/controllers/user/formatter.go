package user

import (
	"be/delivery/controllers/product"
	"time"
)

type UserResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name     string                    `json:"name"`
	Email    string                    `json:"email"`
	Products []product.ProductResponse `json:"products"`
}
