package db

import (
	"log"
	"user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection using GORM
func InitDB() (*gorm.DB, error) {
	dsn := config.GetDBConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	return db, nil
}
