package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
