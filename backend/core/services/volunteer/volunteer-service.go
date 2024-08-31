package Volunteer

import (
	"context"
	"database/sql"
	"hack/core/models"

	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type VolunteerService struct {
	db *sql.DB
}

type VolunteerServiceInterface interface {
	CreateVolunteer(ctx context.Context, business *models.Volunteer) error
	GetVolunteer(ctx context.Context) ([]*models.Volunteer, error)
}

func NewService(db *sql.DB) VolunteerService {
	return VolunteerService{
		db: db,
	}
}

func (s *VolunteerService) CreateVolunteer(ctx context.Context, volunteer *models.Volunteer) error {

	err := volunteer.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting Volunteer into database: %v", err)
		return err
	}
	return nil

}

func (s *VolunteerService) GetVolunteers(ctx context.Context) ([]*models.Volunteer, error) {
	volunteers, err := models.Volunteers().All(ctx, s.db)
	if err != nil {
		log.Printf("Error fetching Volunteers from database: %v", err)
		return nil, err
	}
	return volunteers, nil
}
