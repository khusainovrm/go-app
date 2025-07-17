package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-app/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// GetUsers получает всех пользователей
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User

	if err := h.db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve users",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

// GetUser получает пользователя по ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

// CreateUser создает нового пользователя
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

// UpdateUser обновляет пользователя
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
		return
	}

	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to find user",
			Error:   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

// DeleteUser удаляет пользователя
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
		return
	}

	if err := h.db.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
