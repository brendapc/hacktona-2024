package shelterController

import (
	"hack/core/models"
	service "hack/core/services/shelter"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShelterController struct {
	shelterService service.ShelterService
}

func (c *ShelterController) CreateShelter(ctx *gin.Context) {
	var shelter models.Shelter

	if err := ctx.BindJSON(&shelter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.shelterService.CreateShelter(ctx, &shelter); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, shelter)
}

func Handler(router *gin.RouterGroup, shelterService service.ShelterService) *ShelterController {
	controller := &ShelterController{
		shelterService: shelterService,
	}

	router.POST("/shelters", controller.CreateShelter)

	return controller
}
