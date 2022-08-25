package main

import (
	"contentSquare/src/internal/handlers"
	"contentSquare/src/internal/repositories"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./env/variables.env")
	if err != nil {
		fmt.Println("Error loading env vars " + err.Error())
	}
	fmt.Println("env vars have been loaded...")

	time.Sleep(5000 * time.Millisecond)

	DBRepo := repositories.NewMySqlRepository()
	err = DBRepo.CreateDataBaseAndTable()
	if err != nil {
		fmt.Println("Error creating DB and table: " + err.Error())
	}

	err = DBRepo.IngestFileData(context.Background(), os.Getenv("FILEPATH"))
	if err != nil {
		fmt.Println("Error loading file: " + err.Error())
	}

	r := gin.Default()
	app := r.Group("/contentsquare")

	handlers.NewAggregateHandler(app, DBRepo)
	handlers.NewHealthHandler(app)

	err = r.Run(":8080")
	if err != nil {
		fmt.Println("Error running server: " + err.Error())
	}
}
