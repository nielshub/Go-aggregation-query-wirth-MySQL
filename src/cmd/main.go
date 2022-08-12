package main

import (
	"contentSquare/src/internal/handlers"
	"contentSquare/src/internal/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../env/variables.env")
	if err != nil {
		fmt.Println("Error loading env vars " + err.Error())
	}
	r := gin.Default()
	app := r.Group("/contentsquare")
	DBRepo := repositories.NewMySqlRepository()
	handlers.NewIngestHandler(app, DBRepo)

	err = r.Run(":8080")
	if err != nil {
		fmt.Println("Error running server " + err.Error())
	}
}
