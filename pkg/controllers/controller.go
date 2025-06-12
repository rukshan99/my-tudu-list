package controllers

import (
	"database/sql"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

// GetTask retrieves a task by its ID from the database and returns it as a JSON response.
func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
    // Extract the ID from the URL parameters
    vars := mux.Vars(r)
    idParam := vars["id"]
    if idParam == "" {
        http.Error(w, "Task ID is required", http.StatusBadRequest)
        return
    }

    // Convert the ID to an integer
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid Task ID", http.StatusBadRequest)
        return
    }

    // Call the repository function to get the task by ID
    task, err := repository.GetTaskByID(tc.DB, id)
    if err != nil {
        http.Error(w, "Failed to retrieve task", http.StatusInternalServerError)
        return
    }

    // If the task is not found, return a 404 response
    if task == nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    // Write the task as a JSON response
    if err := utils.WriteJSONResponse(w, http.StatusOK, task); err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
    }
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
