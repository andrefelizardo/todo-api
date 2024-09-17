package main

import (
	"log"

	"github.com/andrefelizardo/todo-api/internal/configs"
	"github.com/andrefelizardo/todo-api/internal/infrastructure"
	"github.com/andrefelizardo/todo-api/internal/routes"
	"github.com/andrefelizardo/todo-api/internal/utils"
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

	utils.InitJWTSecret(&config.JWT)

	router := routes.SetupRouter(db)

	router.Run()
}