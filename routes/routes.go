package routes

import (
	"go-url-shortner/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:code", handlers.RedirectURL)
}
