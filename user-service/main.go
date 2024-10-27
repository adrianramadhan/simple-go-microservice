package main

import (
	"fmt"
	"user-service/internal/db"
)

func main() {
	_, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")
}
