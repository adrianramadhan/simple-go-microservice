package main

import (
	"log"
	"order-service/internal/db"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	orderRepo := repository.NewOrderRepository(db)
	orderSvc := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderSvc)

	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.GetAllOrders)
	r.GET("/orders/:id", orderHandler.GetOrderById)
	r.Run(":8082")
}
