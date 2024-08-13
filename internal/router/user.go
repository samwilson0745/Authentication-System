package router

import (
	"authsys/internal/handler"

	"github.com/gorilla/mux"
)

func UserRouter(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
	return r
}
