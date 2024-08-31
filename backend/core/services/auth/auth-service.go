package auth

import (
	"context"
	"database/sql"
	"errors"
	loginModel "hack/api/controllers/login/models"
	"hack/core/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *AuthService) Authenticate(ctx context.Context, loginReq *loginModel.LoginRequest) (string, error) {
	shelter, err := s.findShelterByEmail(ctx, loginReq.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(shelter.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: shelter.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) findShelterByEmail(ctx context.Context, email string) (*models.Shelter, error) {
	shelter := &models.Shelter{}
	shelter, err := models.Shelters(models.ShelterWhere.Email.EQ(email)).One(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return shelter, nil
}
