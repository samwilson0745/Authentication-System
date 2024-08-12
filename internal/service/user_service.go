package service

import (
	"authsys/internal/database/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateUser(user *models.User) error {
	result := s.DB.Create(user)
	return result.Error
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	result := s.DB.Find(&users)
	return users, result.Error
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := s.DB.First(&user, id)
	return &user, result.Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	result := s.DB.Save(user)
	return result.Error
}

func (s *UserService) DeleteUser(id uint) error {
	result := s.DB.Delete(&models.User{}, id)
	return result.Error
}
