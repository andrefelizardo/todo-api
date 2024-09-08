package main

import (
	"github.com/andrefelizardo/todo-api/internal/routes"
)

func main() {
	router := routes.SetupRouter()

	router.Run()
}