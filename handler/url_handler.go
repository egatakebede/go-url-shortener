package handlers

import (
	"net/http"

	"go-url-shortener/storage"
	"go-url-shortener/utils"

	"github.com/gin-gonic/gin"
)

type Request struct {
	URL string `json:"url"`
}

func ShortenURL(c *gin.Context) {
	var req Request

	// Parse JSON
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate
	if req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL is required",
		})
		return
	}

	// Generate short code
	shortCode := utils.GenerateShortCode(6)

	// Save mapping
	storage.Save(shortCode, req.URL)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}
