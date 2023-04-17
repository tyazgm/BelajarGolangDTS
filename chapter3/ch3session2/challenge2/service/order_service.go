package service

import (
	"challenge2/helper"
	"challenge2/model"
	"challenge2/repository"
	"fmt"
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

func (os *OrderService) GetList(userID string) ([]model.OrderGetResponse, error) {
	orderData, err := os.orderRepository.Get(userID)

	orderDataResponse := make([]model.OrderGetResponse, 0)
	for _, data := range orderData {
		orderDataResponse = append(orderDataResponse, model.OrderGetResponse{
			ID:     data.ID,
			UserID: data.UserID,
			Price:  data.Price,
		})
	}

	fmt.Println(orderDataResponse)

	return orderDataResponse, err
}
