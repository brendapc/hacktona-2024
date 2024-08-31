package shelter

import (
	"context"
	"database/sql"
	"hack/core/models"

	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
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
		log.Printf("Erro ao inserir o shelter no banco de dados: %v", err)
		return err
	}
	return nil
}
