package api

import (
	"database/sql"
	businessController "hack/api/controllers/business"
	loginController "hack/api/controllers/login"
	shelterController "hack/api/controllers/shelter"
	middleware "hack/api/middlewares"
	"hack/core/services/auth"
	"hack/core/services/business"
	"hack/core/services/shelter"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {

	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		jwtKey = []byte("my_secret_key")
	}

	shelterService := shelter.NewService(db)
	authService := auth.NewAuthService(db)
	businessService := business.NewService(db)

	v1 := router.Group("/api/v1")

	loginController.Handler(v1, authService)

	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware(jwtKey))

	shelterController.Handler(v1, shelterService)
	businessController.Handler(v1, businessService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
