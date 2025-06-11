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
