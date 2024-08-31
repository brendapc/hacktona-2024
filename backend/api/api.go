package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Setup(ginApp *gin.Engine, db *sql.DB) {
	// Inicializa os serviços necessários
	// Exemplo:
	// userService := services.NewUserService(db)

	// Configura os controladores e suas respectivas rotas
	// Exemplo:
	// v1 := ginApp.Group("/api/v1")
	// controllers.UserControllerHandler(userService, v1)

	fmt.Println("API OK!")
}
