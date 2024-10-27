package service

import (
	"order-service/internal/dto"
	"order-service/internal/entity"
	"order-service/internal/repository"
)

type OrderService interface {
	CreateOrder(request dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo}
}

func (s *orderService) CreateOrder(request dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	order := entity.Order{
		Product:  request.Product,
		UserID:   request.UserID,
		Quantity: request.Quantity,
	}

	createdOrder, err := s.orderRepo.CreateOrder(&order)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	return dto.CreateOrderResponse{
		ID:       createdOrder.ID,
		Product:  createdOrder.Product,
		UserID:   createdOrder.UserID,
		Quantity: createdOrder.Quantity,
	}, nil
}
