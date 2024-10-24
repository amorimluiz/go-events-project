package repositories

import (
	"amorimluiz/events/db"
	"amorimluiz/events/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func NewEventRepository() *EventRepository {
	return &EventRepository{DB: db.DB}
}

func (r *EventRepository) Create(event *models.Event) (*models.Event, error) {
	event.ID = uuid.New()
	err := r.DB.Omit("User").Create(event).Error

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) FindAll() []models.Event {
	var events []models.Event
	r.DB.Find(&events)
	return events
}

func (r *EventRepository) FindByID(id string) (*models.Event, error) {
	var event models.Event
	err := r.DB.First(&event, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *EventRepository) Update(event *models.Event) (*models.Event, error) {
	err := r.DB.Save(event).Error

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) Delete(id string) error {
	event, err := r.FindByID(id)

	if err != nil {
		return err
	}

	err = r.DB.Delete(event).Error

	if err != nil {
		return err
	}

	return nil
}
