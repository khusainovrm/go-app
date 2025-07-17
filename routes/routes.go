package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-app/handlers"
	"go-app/models"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	// Health check
	// @Summary      Health Check
	// @Description  Проверка состояния сервера
	// @Tags         health
	// @Accept       json
	// @Produce      json
	// @Success      200  {object}  models.HealthResponse
	// @Router       /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.HealthResponse{
			Status:  "ok",
			Message: "Server is running",
		})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// User routes
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
