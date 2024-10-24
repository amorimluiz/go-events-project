package services

import (
	"amorimluiz/events/models"
	"amorimluiz/events/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventService struct {
	repo *repositories.EventRepository
}

func NewEventService() *EventService {
	return &EventService{repo: repositories.NewEventRepository()}
}

func (s *EventService) CreateEvent(event *models.Event) (int, interface{}) {
	event, err := s.repo.Create(event)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Could not create event."}
	}

	return http.StatusCreated, event
}

func (s *EventService) ListEvents() (int, gin.H) {
	return http.StatusOK, gin.H{"data": s.repo.FindAll()}
}

func (s *EventService) GetEvent(id string) (int, interface{}) {
	event, err := s.repo.FindByID(id)

	if err == nil {
		return http.StatusNotFound, gin.H{"message": "Event not found."}
	}

	return http.StatusOK, event
}

func (s *EventService) UpdateEvent(event *models.Event, userId *uuid.UUID) (int, interface{}) {
	event, err := s.repo.Update(event)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Could not update event."}
	}

	return http.StatusOK, event
}

func (s *EventService) DeleteEvent(id string) (int, interface{}) {
	err := s.repo.Delete(id)

	if err != nil {
		return http.StatusInternalServerError, gin.H{"message": "Could not delete event."}
	}

	return http.StatusNoContent, nil
}
