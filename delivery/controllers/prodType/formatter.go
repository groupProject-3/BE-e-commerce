package prodType

import "time"

type ProductTypeRequest struct {
	Name string `json:"name"`
}

type ProductTypeResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string `json:"name"`
}
