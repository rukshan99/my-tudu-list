package controllers

import (
	"database/sql"
	"math/rand"
	"net/http"

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

    // Validate and convert the ID using the util function
    id, err := utils.ValidateAndConvertID(idParam)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
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

// Validate that the description is not empty
// The description field is mandatory and must be provided in the request payload.
// If the description is missing, a 400 Bad Request response is returned.

// Check if priority is provided, else assign default value
// If the priority field is not provided, it is assigned a default value of 3.
// Priority must be between 1 and 3. If the value is outside this range,
// a 400 Bad Request response is returned.
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

// ModifyTask handles the modification of an existing task.
// It expects a JSON payload with the modified task details and validates the input.

// Validate the description field
// The status field must be one of the allowed values: "pending", "in progress", or "completed".
// If the status is invalid, a 400 Bad Request response is returned.

// Validate the priority field
// The priority field must be provided and must be one of the values: 1, 2, or 3.
// If the priority is missing, zero, or outside the allowed range,
// a 400 Bad Request response is returned.
func (tc *TaskController) ModifyTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task

	// Extract the ID from the URL parameters
    vars := mux.Vars(r)
    idParam := vars["id"]

    // Validate and convert the ID using the util function
    id, err := utils.ValidateAndConvertID(idParam)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Parse the JSON payload from the request
    if err := utils.ParseJSONRequest(r, &task); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	// Assign the extracted ID to the task
    task.ID = id

    // Validate the description field
    validDescriptions := map[string]bool{
        "pending":      true,
        "in progress":  true,
        "completed":    true,
    }
    if !validDescriptions[task.Status] {
        http.Error(w, "Invalid task status. Allowed values are: pending, in progress, completed", http.StatusBadRequest)
        return
    }

    // Validate the priority field
    if task.Priority < 1 || task.Priority > 3 {
        http.Error(w, "Task priority must be 1, 2, or 3", http.StatusBadRequest)
        return
    }

    // Call the repository function to update the task
    if err := repository.UpdateTask(tc.DB, task); err != nil {
        http.Error(w, "Failed to update task in database", http.StatusInternalServerError)
        return
    }

    // Write the updated task as a JSON response
    if err := utils.WriteJSONResponse(w, http.StatusOK, task); err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
    }
}

// DeleteTask handles the deletion of a task by its ID.
func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL parameters
	vars := mux.Vars(r)
	idParam := vars["id"]

	// Validate and convert the ID using the util function
    id, err := utils.ValidateAndConvertID(idParam)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Call the repository function to delete the task by ID
	if err := repository.DeleteTask(tc.DB, id); err != nil {
		http.Error(w, "Failed to delete task from database", http.StatusInternalServerError)
		return
	}

	// Write a success response
	w.WriteHeader(http.StatusNoContent)
}
