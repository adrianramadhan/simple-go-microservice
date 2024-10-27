package repository

import (
	"order-service/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *entity.Order) (*entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
	GetOrderById(id uint) (*entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *entity.Order) (*entity.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) GetAllOrders() ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderById(id uint) (*entity.Order, error) {
	var order entity.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
