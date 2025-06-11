package models

type Task struct {
	ID          int    `json:"id"`          // Unique identifier for the task
	Description string `json:"description"` // Task description (mandatory)
	Status      string `json:"status"`      // Task status (e.g., "pending", "completed")
	Priority    int    `json:"priority"`    // Task priority (integer between 1 and 3)
}
