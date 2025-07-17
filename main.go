package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-app/config"
	"go-app/docs"
	"go-app/handlers"
	"go-app/middleware"
	"go-app/models"
	"go-app/routes"
)

// @title           Go API Project
// @version         1.0
// @description     REST API для управления пользователями с использованием Go, Gin, GORM и PostgreSQL
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Инициализация Swagger docs
	docs.SwaggerInfo.Title = "Go API Project"
	docs.SwaggerInfo.Description = "REST API для управления пользователями"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	// Swagger документация
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	log.Printf("Swagger documentation available at http://localhost:%s/swagger/index.html", port)
	log.Fatal(r.Run(":" + port))
}
