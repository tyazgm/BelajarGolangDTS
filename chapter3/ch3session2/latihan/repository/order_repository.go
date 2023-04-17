package repository

import (
	"latihan/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (or *OrderRepository) Add(newOrder model.Order) error {
	tx := or.db.Create(&newOrder)
	return tx.Error
}
