package repository

import (
	// "fmt"
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

func (or *OrderRepository) Get(userID string) ([]model.Order, error) {
	orderData := make([]model.Order, 0)
	// fmt.Println("userID", userID)
	tx := or.db.Where("user_id = ?", userID).Find(&orderData)
	// fmt.Println(orderData)
	return orderData, tx.Error
}
