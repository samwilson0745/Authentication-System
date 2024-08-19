package main

import (
	"authsys/internal/config"
	"authsys/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	db, error := config.InitDB()
	fmt.Print(error)

	r := router.MainRouter(db)

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
