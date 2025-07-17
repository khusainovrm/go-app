package models

import (
	"gorm.io/gorm"
	"time"
)

// User представляет модель пользователя
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Name      string    `json:"name" gorm:"not null" example:"John Doe"`
	Email     string    `json:"email" gorm:"unique;not null" example:"john@example.com"`
	Age       int       `json:"age" example:"30"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T12:00:00Z"`
}

// CreateUserRequest представляет запрос на создание пользователя
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"john@example.com"`
	Age   int    `json:"age" binding:"min=1,max=120" example:"30"`
}

// UpdateUserRequest представляет запрос на обновление пользователя
type UpdateUserRequest struct {
	Name  string `json:"name" example:"Jane Doe"`
	Email string `json:"email" binding:"omitempty,email" example:"jane@example.com"`
	Age   int    `json:"age" binding:"omitempty,min=1,max=120" example:"25"`
}

// Migrate выполняет миграции для всех моделей
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
