package router

import (
	"authsys/internal/handler"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router, authHandler *handler.AuthHandler) {
	r.HandleFunc("/sign-in", authHandler.SignIn).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
}
