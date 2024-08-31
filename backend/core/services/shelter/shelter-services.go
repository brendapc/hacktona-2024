package shelter

import (
	"context"
	"database/sql"
	"errors"
	"hack/core/models"

	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type ShelterService struct {
	db *sql.DB
}

type ShelterServiceInterface interface {
	CreateShelter(ctx context.Context, shelter *models.Shelter) error
}

func NewService(db *sql.DB) ShelterService {
	return ShelterService{
		db: db,
	}
}

func (s *ShelterService) CreateShelter(ctx context.Context, shelter *models.Shelter) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(shelter.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao gerar hash da senha: %v", err)
		return err
	}
	shelter.PasswordHash = string(hashedPassword)

	err = shelter.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting shelter into database: %v", err)
		return err
	}
	return nil
}

func (s *ShelterService) CreateShelterNeed(ctx context.Context, shelter *models.Shelter, shelterNeed *models.ShelterNeed) (*models.ShelterNeed, error) {
	err := shelterNeed.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting shelter need into database: %v", err)
		return nil, err
	}
	return shelterNeed, nil
}

func (s *ShelterService) FindShelterNeedByShelterID(ctx context.Context, shelterID int) (*models.ShelterNeed, error) {
	shelterNeed, err := models.ShelterNeeds(qm.Where("shelter_id = ?", shelterID)).One(ctx, s.db)
	if err != nil {
		log.Printf("Error fetching shelter need from database: %v", err)
		return nil, err
	}
	return shelterNeed, nil

}

func (s *ShelterService) UpdateShelterNeeds(ctx context.Context, shelterNeed *models.ShelterNeed) (*models.ShelterNeed, error) {
	_, err := shelterNeed.Update(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error updating shelter need: %v", err)
		return nil, err
	}
	return shelterNeed, nil
}

func (s *ShelterService) GetShelters(ctx context.Context) ([]*models.Shelter, error) {
	shelters, err := models.Shelters().All(ctx, s.db)
	if err != nil {
		log.Printf("Error fetching shelters from database: %v", err)
		return nil, err
	}
	return shelters, nil
}

func (s *ShelterService) GetShelter(ctx context.Context, id int) (*models.Shelter, error) {
	shelter, err := models.FindShelter(ctx, s.db, id)
	if err != nil {
		log.Printf("Error fetching shelter from database: %v", err)
		return nil, err
	}
	return shelter, nil
}

func (s *ShelterService) AddResident(ctx context.Context, shelter *models.Shelter, resident *models.ShelterResident) (*models.ShelterResident, error) {
	err := resident.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting resident into database: %v", err)
		return nil, err
	}

	shelter.CurrentOccupancy++

	_, err = shelter.Update(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error updating shelter occupancy: %v", err)
		return nil, err
	}

	return resident, nil
}

func (s *ShelterService) RemoveResident(ctx context.Context, shelter *models.Shelter, resident *models.ShelterResident) (*models.ShelterResident, error) {
	_, err := resident.Delete(ctx, s.db)
	if err != nil {
		log.Printf("Error deleting resident from database: %v", err)
		return nil, err
	}

	shelter.CurrentOccupancy--

	_, err = shelter.Update(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error updating shelter occupancy: %v", err)
		return nil, err
	}

	return resident, nil
}

func (s *ShelterService) FindSheltersByNeed(ctx context.Context, item string) ([]*models.Shelter, error) {
	validColumns := map[string]string{
		"bedding_item":        models.ShelterNeedColumns.BeddingItem,
		"food_non_perishable": models.ShelterNeedColumns.FoodNonPerishable,
		"food_perishable":     models.ShelterNeedColumns.FoodPerishable,
		"hygiene_products":    models.ShelterNeedColumns.HygieneProducts,
		"clothing_male":       models.ShelterNeedColumns.ClothingMale,
		"clothing_female":     models.ShelterNeedColumns.ClothingFemale,
		"clothing_children":   models.ShelterNeedColumns.ClothingChildren,
		"medical_supplies":    models.ShelterNeedColumns.MedicalSupplies,
		"pet_food_dogs":       models.ShelterNeedColumns.PetFoodDogs,
		"pet_food_cats":       models.ShelterNeedColumns.PetFoodCats,
		"cleaning_supplies":   models.ShelterNeedColumns.CleaningSupplies,
	}

	column, ok := validColumns[item]
	if !ok {
		return nil, errors.New("invalid search item")
	}

	println("Filtrando pela coluna:", column)

	shelters, err := models.Shelters(
		qm.InnerJoin("shelter_needs ON shelter.id = shelter_needs.shelter_id"),
		qm.Where(column+" = ?", true),
	).All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	println("Número de abrigos encontrados:", len(shelters))

	return shelters, nil
}
