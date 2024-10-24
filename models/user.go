package models

import (
	"amorimluiz/events/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primary_key" json:"id"`
	Email    string    `binding:"required" gorm:"not null;unique" json:"email"`
	Password string    `binding:"required" gorm:"not null" json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = utils.HashPassword(u.Password)

	return err
}
