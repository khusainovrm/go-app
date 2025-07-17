package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-app/config"
	"go-app/handlers"
	"go-app/middleware"
	"go-app/models"
	"go-app/routes"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Инициализация базы данных
	db := config.InitDB()

	// Миграции
	if err := models.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Создание Gin router
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// Инициализация handlers
	userHandler := handlers.NewUserHandler(db)

	// Настройка маршрутов
	routes.SetupRoutes(r, userHandler)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
