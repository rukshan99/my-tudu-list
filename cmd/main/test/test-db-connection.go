package main

import (
	"log"
	
	"my-tudu-list/pkg/config"
)

func main() {
	db, err := config.ConnectToOracle()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
		return
	}
	defer db.Close()

	log.Println("Database connection test successful!")
}
