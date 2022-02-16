package templates

import "time"

type ProductTypeRequest struct {
	Name string `json:"name"`
}

type ProductTypeResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name"`
}