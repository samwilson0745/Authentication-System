package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required,min=3,max=50"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
