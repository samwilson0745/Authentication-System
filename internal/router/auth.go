package router

import (
	"authsys/internal/handler"

	"github.com/gorilla/mux"
)

func AuthRouter(authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/sign-in", authHandler.SignIn).Methods("POST")
	return r
}
