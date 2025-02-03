package controllers

import (
	"api_tracker/config"
	"api_tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TrackRequest(c *gin.Context) {
	apiKey := c.GetHeader("api_key")

	var client models.Client
	if err := config.DB.Where("api_key = ?", apiKey).First(&client).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid API Key"})
		return
	}

	requestLog := models.RequestLog{
		ClientID:  client.ID,
		Endpoint:  c.Request.URL.Path,
		Timestamp: time.Now(),
	}

	var count int
	windowStart := time.Now().Add(-1 * time.Hour)
	config.DB.Model(&models.RequestLog{}).Where("client_id = ? AND timestamp > ?", client.ID, windowStart).Count(&count)
	count++
	if count > client.RateLimit {
		c.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
		return
	} else {
		if err := config.DB.Create(&requestLog).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log request"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successful"})
}
