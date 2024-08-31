package shelterController

import (
	"fmt"
	shelterModels "hack/api/controllers/shelter/models"
	"hack/core/models"
	service "hack/core/services/shelter"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShelterController struct {
	shelterService service.ShelterService
}

func (c *ShelterController) CreateShelter(ctx *gin.Context) {
	var shelter models.Shelter
	var shelterNeeds models.ShelterNeed

	if err := ctx.BindJSON(&shelter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.shelterService.CreateShelter(ctx, &shelter); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	shelterNeeds.ShelterID = shelter.ID

	if _, err := c.shelterService.CreateShelterNeed(ctx, &shelter, &shelterNeeds); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var shelterPresenter shelterModels.ShelterPresenter

	shelterPresenter.InitFromModel(shelter)

	ctx.JSON(http.StatusCreated, shelterPresenter)
}

func (c *ShelterController) GetSheltersByNeed(ctx *gin.Context) {
	searchItem := ctx.Query("search")

	if searchItem == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Debug: Verificar o valor do searchItem
	println("Item de busca:", searchItem)

	shelters, err := c.shelterService.FindSheltersByNeed(ctx, searchItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, shelters)
}

func (c *ShelterController) UpdateShelterNeeds(ctx *gin.Context) {
	shelterID := ctx.Request.Header.Get("shelter_id")
	shelterIDInt, _ := strconv.Atoi(shelterID)
	fmt.Println(shelterIDInt)

	var shelterNeeds *models.ShelterNeed

	shelterNeeds, err := c.shelterService.FindShelterNeedByShelterID(ctx, shelterIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("needs", shelterNeeds)
	if err := ctx.BindJSON(&shelterNeeds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("needs2", shelterNeeds)
	shelterNeeds, err = c.shelterService.UpdateShelterNeeds(ctx, shelterNeeds)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, shelterNeeds)
}

func (c *ShelterController) GetShelters(ctx *gin.Context) {
	shelters, err := c.shelterService.GetShelters(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var shelterPresenters []shelterModels.ShelterPresenter

	for _, shelter := range shelters {
		var shelterPresenter shelterModels.ShelterPresenter
		shelterPresenter.InitFromModel(*shelter)
		shelterPresenters = append(shelterPresenters, shelterPresenter)
	}

	ctx.JSON(http.StatusOK, shelterPresenters)
}

func (c *ShelterController) AddResident(ctx *gin.Context) {
	shelterID := ctx.Request.Header.Get("shelter_id")

	var resident *models.ShelterResident
	if err := ctx.BindJSON(&resident); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelterIDInt, _ := strconv.Atoi(shelterID)

	shelter, err := c.shelterService.GetShelter(ctx, shelterIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resident.ShelterID = shelter.ID

	resident, err = c.shelterService.AddResident(ctx, shelter, resident)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resident)
}

func (c *ShelterController) RemoveResident(ctx *gin.Context) {
	shelterID := ctx.Request.Header.Get("shelter_id")

	var resident *models.ShelterResident
	if err := ctx.BindJSON(&resident); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelterIDInt, _ := strconv.Atoi(shelterID)

	shelter, err := c.shelterService.GetShelter(ctx, shelterIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resident.ShelterID = shelter.ID

	resident, err = c.shelterService.RemoveResident(ctx, shelter, resident)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resident)
}

func (c *ShelterController) GetShelter(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	shelter, err := c.shelterService.GetShelter(ctx, idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var shelterPresenter shelterModels.ShelterPresenter
	shelterPresenter.InitFromModel(*shelter)

	ctx.JSON(http.StatusOK, shelterPresenter)
}

func Handler(router *gin.RouterGroup, shelterService service.ShelterService) *ShelterController {
	controller := &ShelterController{
		shelterService: shelterService,
	}

	router.POST("/shelters", controller.CreateShelter)
	router.GET("/shelters", controller.GetShelters)
	router.GET("/shelters/:id", controller.GetShelter)
	router.POST("/shelters/residents", controller.AddResident)
	router.DELETE("/shelters/residents", controller.RemoveResident)
	router.PUT("/shelters/needs", controller.UpdateShelterNeeds)
	router.GET("/shelters/needs", controller.GetSheltersByNeed)

	return controller
}
