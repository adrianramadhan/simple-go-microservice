package service

import (
	"order-service/internal/dto"
	"order-service/internal/entity"
	"order-service/internal/repository"
)

type OrderService interface {
	CreateOrder(request dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
	GetAllOrders() ([]dto.CreateOrderResponse, error)
	GetOrderById(id uint) (dto.CreateOrderResponse, error)
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

func (s *orderService) GetAllOrders() ([]dto.CreateOrderResponse, error) {
	orders, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var response []dto.CreateOrderResponse
	for _, order := range orders {
		response = append(response, dto.CreateOrderResponse{
			ID:       order.ID,
			Product:  order.Product,
			UserID:   order.UserID,
			Quantity: order.Quantity,
		})
	}

	return response, nil
}

func (s *orderService) GetOrderById(id uint) (dto.CreateOrderResponse, error) {
	order, err := s.orderRepo.GetOrderById(id)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	return dto.CreateOrderResponse{
		ID:       order.ID,
		Product:  order.Product,
		UserID:   order.UserID,
		Quantity: order.Quantity,
	}, nil
}
