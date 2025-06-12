# My TuDu List

## Overview
Golang based To-Do list application with Oracle database integration. This application provides CRUD operations for managing tasks.

## Prerequisites
- Go installed
- Oracle Instant Client
- Postman or curl for testing

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/my-tudu-list.git
   cd my-tudu-list
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment variables:
   - Set the `ORACLE_DSN` environment variable to your Oracle database connection string.
     Example:
     ```bash
     user/password@host:port/service_name
     ```
     For Windows (PowerShell)
        ```bash
        $env:ORACLE_DSN="user/password@host:port/service_name"
        ```
     For Windows (Command Prompt)
        ```bash
        set ORACLE_DSN="user/password@host:port/service_name"
        ```
     For Linux/macOS
        ```bash
        export ORACLE_DSN="user/password@host:port/service_name"
        ```

4. Run the application:
   ```bash
   go run cmd/main/main.go
   ```
    The server will start on http://localhost:8080.

## Routes
### Create Task
- **Endpoint**: `POST /api/tasks`
- **Payload**:
  ```json
  {
      "description": "Task description",
      "priority": 1
  }
  ```
- **Validations**:
  - `description` must not be empty.
  - Initial `status` will always be defaulted to `"pending"`.
  - `priority` must be between `1` and `3`.

### Get Task
- **Endpoint**: `GET /api/tasks/{id}`
- **Response**:
  ```json
  {
      "id": 1,
      "description": "Task description",
      "status": "pending",
      "priority": 1
  }
  ```

### Get All Tasks
- **Endpoint**: `GET /api/tasks`
- **Response**:
  ```json
  [
      {
          "id": 1,
          "description": "Task description",
          "status": "pending",
          "priority": 1
      },
      {
          "id": 2,
          "description": "Another task",
          "status": "completed",
          "priority": 2
      }
  ]
  ```

### Modify Task
- **Endpoint**: `PUT /api/tasks/{id}`
- **Payload**:
  ```json
  {
      "description": "Updated description",
      "status": "completed",
      "priority": 2
  }
  ```
- **Validations**:
  - `description` must not be empty.
  - `status` must be one of `"pending"`, `"in progress"`, or `"completed"`.
  - `priority` must be between `1` and `3`.

### Delete Task
- **Endpoint**: `DELETE /api/tasks/{id}`

## Environment Variables
- `ORACLE_DSN`: Oracle database connection string.

## Testing
- Use Postman or `curl` to test the routes.
- Import the provided Postman collection (`postman/my-tudu-list.postman_collection.json`) for easier testing.
- Example `curl` command for creating a task:
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"description": "New Task", "status": "pending", "priority": 1}' http://localhost:8080/tasks
  ```

## Project Structure
```
my-tudu-list/
├── cmd/
│   ├── main/
│   │   ├── main.go                # Entry point of the application
│   │   ├── test-db-connection.go  # Script to test database connection
├── pkg/
│   ├── config/
│   │   ├── app.go                 # Contains database connection logic
│   ├── controllers/
│   │   ├── controller.go          # Handles HTTP requests and business logic
│   ├── models/
│   │   ├── task.go                # Defines the Task model
│   ├── repository/
│   │   ├── repository.go          # Handles database operations
│   ├── routes/
│   │   ├── route.go               # Defines API routes
│   ├── utils/
│   │   ├── util.go                # Utility functions for JSON handling and ID validation
├── postman/
│   ├── my-tudu-list.postman_collection.json # Postman collection for API testing
├── README.md                      # Documentation for the project
├── go.mod                         # Go module file
├── go.sum                         # Go dependencies checksum file
```

### Explanation
- **cmd/main/main.go**: The main entry point of the application where the server is initialized and started.
- **cmd/main/test-db-connection.go**: A utility script to test the Oracle database connection.
- **pkg/config/app.go**: Contains the logic for connecting to the Oracle database using environment variables.
- **pkg/controllers/controller.go**: Implements the controller functions for handling CRUD operations.
- **pkg/models/task.go**: Defines the `Task` struct used across the application.
- **pkg/repository/repository.go**: Contains functions for interacting with the database (e.g., insert, update, delete).
- **pkg/routes/route.go**: Maps API endpoints to their respective controller functions.
- **pkg/utils/util.go**: Provides utility functions for JSON response handling and ID validation.
- **postman/my-tudu-list.postman_collection.json**: A Postman collection for testing the API endpoints.
- **go.mod**: Specifies the Go module and its dependencies.
- **go.sum**: Contains checksums for the module dependencies.

