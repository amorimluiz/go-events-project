package repositories

import (
	"amorimluiz/events/db"
	"amorimluiz/events/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegistrationRepository struct {
	DB *gorm.DB
}

func NewRegistrationRepository() *RegistrationRepository {
	return &RegistrationRepository{DB: db.DB}
}

func (r *RegistrationRepository) RegisterUserInEvent(userID, eventID uuid.UUID) (*models.Registration, error) {
	registration := models.NewRegistration(userID, eventID)

	err := r.DB.Create(registration).Error

	if err != nil {
		return nil, err
	}

	return registration, nil
}

func (r *RegistrationRepository) UnregisterUserFromEvent(userID, eventID uuid.UUID) error {
	err := r.DB.Delete(&models.Registration{}, "user_id = ? AND event_id = ?", userID, eventID).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationRepository) FindByUserIDAndEventID(userID, eventID uuid.UUID) (*models.Registration, error) {
	var registration models.Registration
	err := r.DB.First(&registration, "user_id = ? AND event_id = ?", userID, eventID).Error

	if err != nil {
		return nil, err
	}

	return &registration, nil
}
