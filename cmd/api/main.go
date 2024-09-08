package main

import "github.com/andrefelizardo/todo-api/configs"

func main() {
	router := configs.SetupRouter()

	router.Run()
}