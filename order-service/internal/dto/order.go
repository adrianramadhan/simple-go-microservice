package dto

import "time"

type CreateOrderRequest struct {
	Product  string `json:"product" binding:"required"`
	UserID   uint   `json:"user_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type CreateOrderResponse struct {
	ID        uint      `json:"id"`
	Product   string    `json:"product"`
	UserID    uint      `json:"user_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
