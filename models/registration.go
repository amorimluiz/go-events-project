package models

import "github.com/google/uuid"

type Registration struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	EventID uuid.UUID `gorm:"type:char(36);not null" json:"eventId"`
	User    *User     `binding:"-" json:"user,omitempty"`
	Event   *Event    `binding:"-" json:"event,omitempty"`
}

func NewRegistration(userID, eventID uuid.UUID) *Registration {
	return &Registration{
		ID:      uuid.New(),
		UserID:  userID,
		EventID: eventID,
	}
}
