package api

import (
	"database/sql"
	shelterController "hack/api/controllers/shelter"
	"hack/core/services/shelter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	shelterService := shelter.NewService(db)

	v1 := router.Group("/api/v1")

	shelterController.Handler(v1, shelterService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
