package loginController

import (
	"hack/core/services/auth"
	"net/http"

	loginModel "hack/api/controllers/login/models"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	authService *auth.AuthService
}

func NewLoginController(authService *auth.AuthService) *LoginController {
	return &LoginController{
		authService: authService,
	}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var loginReq loginModel.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := c.authService.Authenticate(ctx, &loginReq)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Handler(router *gin.RouterGroup, authService *auth.AuthService) *LoginController {
	controller := NewLoginController(authService)

	router.POST("/login", controller.Login)

	return controller
}
