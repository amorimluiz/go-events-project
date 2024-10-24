package services

import (
	"amorimluiz/events/repositories"
	"net/http"

	"github.com/google/uuid"
)

type RegistrationService struct {
	rep *repositories.RegistrationRepository
}

func NewRegistrationService() *RegistrationService {
	return &RegistrationService{rep: repositories.NewRegistrationRepository()}
}

func (s *RegistrationService) RegisterUserInEvent(userID, eventID string) (int, interface{}) {
	parsedUserID, parseUserErr := uuid.Parse(userID)
	parsedEventID, parseEventErr := uuid.Parse(eventID)

	eventAlreadyRegistered, _ := s.rep.FindByUserIDAndEventID(parsedUserID, parsedEventID)

	if eventAlreadyRegistered != nil {
		return http.StatusConflict, map[string]string{"message": "User already registered in event."}
	}

	if parseUserErr != nil || parseEventErr != nil {
		return http.StatusBadRequest, map[string]string{"message": "Invalid user or event ID."}
	}

	registration, err := s.rep.RegisterUserInEvent(parsedUserID, parsedEventID)

	if err != nil {
		return http.StatusInternalServerError, map[string]string{"message": "Could not register user in event."}
	}

	return http.StatusCreated, registration
}

func (s *RegistrationService) UnregisterUserFromEvent(userID, eventID string) (int, interface{}) {
	parsedUserID, parseUserErr := uuid.Parse(userID)
	parsedEventID, parseEventErr := uuid.Parse(eventID)

	if parseUserErr != nil || parseEventErr != nil {
		return http.StatusBadRequest, map[string]string{"message": "Invalid user or event ID."}
	}

	err := s.rep.UnregisterUserFromEvent(parsedUserID, parsedEventID)

	if err != nil {
		return http.StatusInternalServerError, map[string]string{"message": "Could not unregister user from event."}
	}

	return http.StatusNoContent, nil
}
