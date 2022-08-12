package handlers

import (
	"contentSquare/src/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func NewHealthHandler(app *gin.RouterGroup) {
	app.GET("/health", health)
}

func health(c *gin.Context) {
	response := models.Health{
		Status:  http.StatusOK,
		Message: "DataStore is up and running",
		Version: os.Getenv("VERSION"),
		Stack:   os.Getenv("ENV"),
	}
	c.JSON(http.StatusOK, response)
}
