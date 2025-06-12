package routes

import (
	"database/sql"
	
	"my-tudu-list/pkg/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the routes for the application
func RegisterRoutes(db *sql.DB, router *mux.Router) {
	taskController := controllers.NewTaskController(db)

	router.HandleFunc("/api/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskController.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks", taskController.CreateTask).Methods("POST")
	//router.HandleFunc("/api/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	//router.HandleFunc("/api/tasks/{id}", taskController.DeleteTask).Methods("DELETE")
}
