package controllers

import (
	"api_tracker/config"
	"api_tracker/dto"
	"api_tracker/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateClient(c *gin.Context) {
	var clientRequest dto.CreateClientRequest

	if err := c.ShouldBindJSON(&clientRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := models.Client{
		Name:      clientRequest.Name,
		Email:     clientRequest.Email,
		CreatedAt: time.Now(),
		APIKey:    uuid.New().String(),
		RateLimit: clientRequest.RateLimit,
		Password:  clientRequest.Password,
	}

	if err := config.DB.Where("email = ?", clientRequest.Email).First(&client).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with such email already exists"})
		return
	}

	if err := config.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}

func GetUser(c *gin.Context) {
	var clientRequest dto.LoginClientRequest

	if err := c.ShouldBindJSON(&clientRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var client models.Client

	if err := config.DB.Where("email = ?", clientRequest.Email).First(&client).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, client)
}
