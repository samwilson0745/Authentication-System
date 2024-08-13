package service

import (
	"authsys/internal/database/models"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (s *AuthService) SignIn(user *models.User) (string, error) {
	result := s.DB.Create(user)
	if result.Error == nil {
		return "User registered", nil
	}
	return "Error while register user", nil
}
