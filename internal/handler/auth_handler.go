package handler

import (
	"authsys/internal/database/models"
	"authsys/internal/service"
	"authsys/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var message string
	validate := validator.New()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := validate.Struct(user)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			validationErrors[field] = fmt.Sprintf("Field '%s' failed validation for tag '%s'", field, tag)
		}
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validation error",
			"details": validationErrors,
		})
		return
	}
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	message, err = h.Service.SignIn(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
	})
}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := validate.Struct(&credentials)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			validationErrors[field] = fmt.Sprintf("Field '%s' failed validation for tag '%s'", field, tag)
		}
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Validation error",
			"details": validationErrors,
		})
		return
	}

	message, err := h.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, message, http.StatusUnauthorized)
		return
	}

	// Return success response
	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
