package businessController

import (
	"hack/core/models"
	service "hack/core/services/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	businessService service.BusinessService
}

func (c *BusinessController) CreateBusiness(ctx *gin.Context) {
	var business models.Business

	if err := ctx.BindJSON(&business); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.businessService.CreateBusiness(ctx, &business); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, business)
}

func (c *BusinessController) GetBusinesses(ctx *gin.Context) {
	businesses, err := c.businessService.GetBusinesses(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, businesses)
}

func Handler(router *gin.RouterGroup, businessService service.BusinessService) *BusinessController {
	controller := &BusinessController{
		businessService: businessService,
	}

	router.POST("/business", controller.CreateBusiness)
	router.GET("/business", controller.GetBusinesses)

	return controller
}
