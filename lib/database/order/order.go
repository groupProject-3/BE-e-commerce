package order

import (
	"be/models"

	"gorm.io/gorm"
)

type OrderDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *OrderDb  {
	return &OrderDb{
		db: db,
	}
}

func (od *OrderDb) Create(user_id uint, newOrder models.Order) (models.Order, error) {
	newOrder.User_id = user_id
	if err := od.db.Create(&newOrder).Error ; err != nil {
		return models.Order{}, err
	}
	return newOrder, nil
}

func (od *OrderDb) UpdateById()