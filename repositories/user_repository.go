package repositories

import (
	"amorimluiz/events/db"
	"amorimluiz/events/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: db.DB}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	user.ID = uuid.New()
	err := r.DB.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
