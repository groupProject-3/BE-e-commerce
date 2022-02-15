package user

import (
	"time"
)

type UserResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
