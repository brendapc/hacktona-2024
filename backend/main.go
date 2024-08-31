package main

import (
	"hack/api"
	middleware "hack/api/middlewares"
	"hack/infrastructure/database"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	api.Setup(r, db)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
