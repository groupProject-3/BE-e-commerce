package user

import (
	"time"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name  string `json:"name"`
	Email string `json:"email"`
}


