package service

import (
	"authsys/internal/database/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
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
func (s *AuthService) Login(email string, password string) (string, error) {
	var user models.User
	// Find the user by email
	result := s.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return "User not found", errors.New("user not found")
		}
		return "Error querying user", result.Error
	}

	// Compare the provided password with the hashed password stored in the database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "Invalid credentials", errors.New("invalid credentials")
	}

	return "Login successful", nil
}
