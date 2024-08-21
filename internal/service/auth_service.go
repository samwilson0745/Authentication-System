package service

import (
	"authsys/internal/database/models"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	token, err := generateJWT(user.Email, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

/*
Function to generate the JWT token for the user
*/
func generateJWT(email string, password string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtKey, err := base64.StdEncoding.DecodeString(os.Getenv("JWT_KEY"))
	if err != nil {
		log.Fatalf("Failed to decode JWT key: %v", err)
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
