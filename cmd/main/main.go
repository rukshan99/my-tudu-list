package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"my-tudu-list/pkg/config"
	"my-tudu-list/pkg/repository"
	"my-tudu-list/pkg/routes"
)

func main() {
	// Initialize the database connection
	db, err := config.ConnectToOracle()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	// Create the tasks table if it doesn't exist
	repository.CreateTasksTable(db)

	// Register routes and pass the database connection
	router := mux.NewRouter()
	routes.RegisterRoutes(db, router)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
