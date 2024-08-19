package router

import (
	"authsys/internal/handler"
	"authsys/internal/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func MainRouter(DB *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	authService := &service.AuthService{DB: DB}
	authHandler := &handler.AuthHandler{Service: authService}

	userService := &service.UserService{DB: DB}
	userHandler := &handler.UserHandler{Service: userService}

	// Attach the /auth routes
	authSubRouter := r.PathPrefix("/auth").Subrouter()
	AuthRouter(authSubRouter, authHandler)

	// Attach the /user routes
	userSubRouter := r.PathPrefix("/user").Subrouter()
	UserRouter(userSubRouter, userHandler)

	return r
}
