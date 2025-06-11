package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror"
)

// ConnectToOracle establishes a connection to the Oracle database using the godror driver.
// It returns a pointer to the sql.DB object and an error if any occurs during the connection process.
func ConnectToOracle() (*sql.DB, error) {
	//ORACLE_DSN => "username/password@host:port/sid"
    dsn := os.Getenv("ORACLE_DSN")

	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Oracle database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping Oracle database: %v", err)
		return nil, err
	}

	fmt.Println("Successfully connected to Oracle database!")
	return db, nil
}
