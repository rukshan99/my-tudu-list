package controllers

import (
	"database/sql"
	"math/rand"
	"net/http"

	"my-tudu-list/pkg/models"
	"my-tudu-list/pkg/repository"
	"my-tudu-list/pkg/utils"
)

type TaskController struct {
	DB *sql.DB
}

// Constructor for TaskController
func NewTaskController(db *sql.DB) *TaskController {
	return &TaskController{DB: db}
}

// CreateTask handles the creation of a new task
// It expects a JSON payload with the task details and validates the input.
func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := utils.ParseJSONRequest(r, &task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate that the description is not empty
	if task.Description == "" {
		http.Error(w, "Task description is required", http.StatusBadRequest)
		return
	}

	// Check if priority is provided, else assign default value
	if task.Priority == 0 {
		task.Priority = 3
	} else {
		// Validate that the priority is between 1 and 3
		if task.Priority < 1 || task.Priority > 3 {
			http.Error(w, "Task priority must be between 1 and 3", http.StatusBadRequest)
			return
		}
	}

	task.ID = rand.Intn(1000)
	task.Status = "pending"

	// Insert the task into the database
	if err := repository.InsertTask(tc.DB, task); err != nil {
		http.Error(w, "Failed to insert task into database", http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSONResponse(w, http.StatusCreated, task); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

// GetTasks retrieves all tasks from the database and returns them as a JSON response.
func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks(tc.DB)
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSONResponse(w, http.StatusOK, tasks); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
