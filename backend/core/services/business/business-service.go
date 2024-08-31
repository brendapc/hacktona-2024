package business

import (
	"context"
	"database/sql"
	"hack/core/models"

	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BusinessService struct {
	db *sql.DB
}

type BusinessServiceInterface interface {
	CreateBusiness(ctx context.Context, business *models.Business) error
	GetBusinesses(ctx context.Context) ([]*models.Business, error)
}

func NewService(db *sql.DB) BusinessService {
	return BusinessService{
		db: db,
	}
}

func (s *BusinessService) CreateBusiness(ctx context.Context, business *models.Business) error {

	err := business.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting business into database: %v", err)
		return err
	}
	return nil

}

func (s *BusinessService) GetBusinesses(ctx context.Context) ([]*models.Business, error) {
	businesses, err := models.Businesses().All(ctx, s.db)
	if err != nil {
		log.Printf("Error fetching businesses from database: %v", err)
		return nil, err
	}
	return businesses, nil
}
