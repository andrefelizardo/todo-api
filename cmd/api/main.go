package main

import (
	"database/sql"
	"log"

	router "github.com/andrefelizardo/todo-api/internal"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	router := router.SetupRouter()

	router.Run()

}