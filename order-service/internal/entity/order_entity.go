package entity

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Product   string    `gorm:"size:255;not null"`
	UserID    uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
