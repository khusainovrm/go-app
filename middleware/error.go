package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-app/models"
)

func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Success: false,
				Message: "Internal server error",
				Error:   err,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
