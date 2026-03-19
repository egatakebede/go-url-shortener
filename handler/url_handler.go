package handlers

import (
	"net/http"
	"net/url"

	"go-url-shortener/storage"
	"go-url-shortener/utils"

	"github.com/gin-gonic/gin"
)

// Request struct for /shorten
type Request struct {
	URL string `json:"url" binding:"required"`
}

// ShortenURL handles POST /shorten
func ShortenURL(c *gin.Context) {
	var req Request

	// Parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate URL format
	if _, err := url.ParseRequestURI(req.URL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL",
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

// Redirect handles GET /:code
func Redirect(c *gin.Context) {
	code := c.Param("code") // get short code from URL

	// Lookup original URL in storage
	originalURL, exists := storage.Get(code)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	// Redirect to the original URL
	c.Redirect(http.StatusFound, originalURL)
}
