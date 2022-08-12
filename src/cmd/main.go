package main

import (
	"contentSquare/src/internal/handlers"
	"contentSquare/src/internal/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app := r.Group("/contentsquare")
	DBRepo := repositories.NewMySqlRepository()
	handlers.NewIngestHandler(app, DBRepo)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error running server " + err.Error())
	}
}
