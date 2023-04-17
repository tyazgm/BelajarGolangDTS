package service

import (
	"latihan/helper"
	"latihan/model"
	"latihan/repository"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (os *OrderService) Create(request model.OrderCreateRequest, userID string) (model.OrderCreateResponse, error) {
	id := helper.GenerateID()

	order := model.Order{
		ID:     id,
		UserID: userID,
		Price:  request.Price,
	}

	err := os.orderRepository.Add(order)

	return model.OrderCreateResponse{
		ID:     id,
		UserID: userID,
		Price:  request.Price,
	}, err
}
