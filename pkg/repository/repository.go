package repository

import (
	"database/sql"
	"log"
	
	"my-tudu-list/pkg/models"
)

// CreateTasksTable creates the "rukjlk_task_rtb" table in the database if it doesn't exist.
func CreateTasksTable(db *sql.DB) error {
	query := `
    BEGIN
        EXECUTE IMMEDIATE '
        CREATE TABLE rukjlk_task_rtb (
            id NUMBER PRIMARY KEY,
            description VARCHAR2(255) NOT NULL,
            status VARCHAR2(50) NOT NULL,
            priority NUMBER NOT NULL
        )';
    EXCEPTION
        WHEN OTHERS THEN
            IF SQLCODE = -955 THEN
                NULL; -- Ignore "name already used by an existing object" error
            ELSE
                RAISE;
            END IF;
    END;`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Failed to create rukjlk_task_rtb table: %v", err)
		return err
	}
	log.Println("rukjlk_task_rtb table created successfully!")
	return nil
}

// GetAllTasks retrieves all tasks from the "rukjlk_task_rtb" table.
func GetAllTasks(db *sql.DB) ([]models.Task, error) {
	query := `SELECT id, description, status, priority FROM rukjlk_task_rtb`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to retrieve tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Status, &task.Priority); err != nil {
			log.Printf("Failed to scan task: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetTaskByID retrieves a task from the "rukjlk_task_rtb" table by its ID.
func GetTaskByID(db *sql.DB, id int) (*models.Task, error) {
    query := `SELECT id, description, status, priority FROM rukjlk_task_rtb WHERE id = :1`
    row := db.QueryRow(query, id)

    var task models.Task
    if err := row.Scan(&task.ID, &task.Description, &task.Status, &task.Priority); err != nil {
        if err == sql.ErrNoRows {
            log.Printf("No task found with ID %d", id)
            return nil, nil // Return nil if no rows are found
        }
        log.Printf("Failed to retrieve task with ID %d: %v", id, err)
        return nil, err
    }

    return &task, nil
}

// InsertTask inserts a new task into the "rukjlk_task_rtb" table.
func InsertTask(db *sql.DB, task models.Task) error {
	query := `INSERT INTO rukjlk_task_rtb (id, description, status, priority) VALUES (:1, :2, :3, :4)`
	_, err := db.Exec(query, task.ID, task.Description, task.Status, task.Priority)
	if err != nil {
		log.Printf("Failed to insert task: %v", err)
		return err
	}
	return nil
}

// UpdateTask updates an existing task in the "rukjlk_task_rtb" table.
func UpdateTask(db *sql.DB, task models.Task) error {
	query := `UPDATE rukjlk_task_rtb SET description = :1, status = :2, priority = :3 WHERE id = :4`
	_, err := db.Exec(query, task.Description, task.Status, task.Priority, task.ID)
	if err != nil {
		log.Printf("Failed to update task with ID %d: %v", task.ID, err)
		return err
	}
	return nil
}

// DeleteTask deletes a task from the "rukjlk_task_rtb" table by its ID.
func DeleteTask(db *sql.DB, id int) error {
	query := `DELETE FROM rukjlk_task_rtb WHERE id = :1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete task with ID %d: %v", id, err)
		return err
	}
	log.Printf("Task with ID %d deleted successfully", id)
	return nil
}