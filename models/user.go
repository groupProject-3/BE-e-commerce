package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string    `gorm:"not null;type:varchar(100)"`
	Email    string    `gorm:"unique;index;not null;type:varchar(100)"`
	Password string    `gorm:"unique;not null;type:varchar(100)"`
	Products []Product `gorm:"foreignKey:User_id"`
	Carts    []Cart    `gorm:"foreignKey:User_id"`
	Orders   []Order   `gorm:"foreignKey:User_id"`
}
