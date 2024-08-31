package volunteerController

import (
	"hack/core/models"
	"net/http"

	service "hack/core/services/volunteer"

	"github.com/gin-gonic/gin"
)

type VolunteerController struct {
	volunteerService service.VolunteerService
}

func (c *VolunteerController) CreateVolunteer(ctx *gin.Context) {
	var volunteer models.Volunteer

	if err := ctx.BindJSON(&volunteer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.volunteerService.CreateVolunteer(ctx, &volunteer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, volunteer)
}

func (c *VolunteerController) GetVolunteers(ctx *gin.Context) {
	volunteers, err := c.volunteerService.GetVolunteers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, volunteers)
}

func Handler(router *gin.RouterGroup, volunteerService service.VolunteerService) *VolunteerController {
	controller := &VolunteerController{
		volunteerService: volunteerService,
	}

	router.POST("/volunteer", controller.CreateVolunteer)
	router.GET("/volunteer", controller.GetVolunteers)

	return controller
}
