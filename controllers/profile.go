package controllers

import (
	"net/http"

	"github.com/ar-tiwari75/wealth-manager/models"
	"github.com/ar-tiwari75/wealth-manager/storage"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.User
	if err := storage.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Convert User to DTO(Data Transfer Object)
	profile := models.ProfileResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	// Fetch from DB using userID
	c.JSON(http.StatusOK, profile)
}
