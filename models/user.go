package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Migrate выполняет миграции для всех моделей
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
