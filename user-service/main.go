package main

import (
	"log"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetAllUsers)
	r.GET("/users/:id", userHandler.GetUserById)
	r.Run(":8081")
}
