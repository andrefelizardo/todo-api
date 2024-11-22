package main

import (
	"database/sql"
	"log"

	"github.com/andrefelizardo/todo-api/internal/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	router := routes.SetupRouter()

	router.Run()

}