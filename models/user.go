package models

import (
	"be/delivery/controllers/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `gorm:"not null;type:varchar(100)"`
	Email    string    `gorm:"unique;index;not null;type:varchar(100)"`
	Password string    `gorm:"unique;not null;type:varchar(100)"`
	Products []Product `gorm:"foreignKey:User_ID"`
}

func (u *User) ToUserResponse() user.UserResponse {
	return user.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,

		Name:  u.Name,
		Email: u.Password,
	}
}
