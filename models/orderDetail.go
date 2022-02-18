package models

import (
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model
	Cart_id  uint
	Order_id uint
	Qty      uint
	Price    uint
}
