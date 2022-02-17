package templates

import (
	"be/models"
	"time"
)

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name  string `json:"name"`
	Email string `json:"email"`
}

type GetUserResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    models.User `json:"data"`
}
