package main

import (
	"log"

	"github.com/andrefelizardo/todo-api/internal/configs"
	"github.com/andrefelizardo/todo-api/internal/infrastructure"
	"github.com/andrefelizardo/todo-api/internal/routes"
)

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	db, err := infrastructure.NewDatabase(&config.DB)
	if err != nil {
		log.Fatal("Error loading database config", err)
	}

	router := routes.SetupRouter(db)

	router.Run()
}