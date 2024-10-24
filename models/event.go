package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Name        string    `binding:"required" gorm:"not null" json:"name"`
	Description string    `binding:"required" gorm:"not null" json:"description"`
	Location    string    `binding:"required" gorm:"not null" json:"location"`
	DateTime    time.Time `binding:"required" gorm:"not null" json:"dateTime"`
	UserID      uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	User        *User     `binding:"-" json:"user,omitempty"`
}
