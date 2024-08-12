package main

import (
	"authsys/internal/config"
	"authsys/internal/handler"
	"authsys/internal/router"
	"authsys/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	db, error := config.InitDB()
	fmt.Print(error)
	// Create the user service
	userService := &service.UserService{DB: db}

	// Create the user handler
	userHandler := &handler.UserHandler{Service: userService}

	// Set up the router
	r := router.SetupRouter(userHandler)

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
