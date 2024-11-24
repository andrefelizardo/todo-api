package main

import (
	"database/sql"
	"log"
	"os"

	router "github.com/andrefelizardo/todo-api/internal"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := router.SetupRouter(db)

	router.Run()

}