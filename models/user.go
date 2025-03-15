package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id       uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	name     string    `json:"name"`
	email    string    `json:"email"`
	password string    `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.id = uuid.New()
	return
}
