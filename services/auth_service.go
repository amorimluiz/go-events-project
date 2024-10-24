package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type AuthService struct {
	jwtSecret string
}

var authService *AuthService

func NewAuthService() *AuthService {
	if authService == nil {
		err := godotenv.Load()

		if err != nil {
			panic("Error loading .env file")
		}

		authService = &AuthService{jwtSecret: os.Getenv("JWT_SECRET")}
	}

	return authService
}

func (a *AuthService) GenerateToken(userId uuid.UUID, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"iss":   "events",
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(a.jwtSecret))
}

func (a *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.jwtSecret), nil
	})
}
