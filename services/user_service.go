package services

import (
	"amorimluiz/events/models"
	"amorimluiz/events/repositories"
	"amorimluiz/events/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repositories.NewUserRepository()}
}

func getAuthToken(user *models.User) (int, interface{}) {
	token, err := NewAuthService().GenerateToken(user.ID, user.Email)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Could not generate token."}
	}

	return http.StatusCreated, gin.H{"authToken": token}
}

func (s *UserService) SignUp(user *models.User) (int, interface{}) {
	_, err := s.repo.Create(user)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Could not create user."}
	}

	return getAuthToken(user)
}

func (s *UserService) Login(loginPayload *models.User) (int, interface{}) {
	user, err := s.repo.FindByEmail(loginPayload.Email)

	if err != nil {
		return http.StatusNotFound, gin.H{"message": "User not found."}
	}

	if !utils.ValidatePassword(user.Password, loginPayload.Password) {
		return http.StatusUnauthorized, gin.H{"message": "Invalid password."}
	}

	return getAuthToken(user)
}
