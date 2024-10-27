package repository

import (
	"order-service/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *entity.Order) (*entity.Order, error)
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
