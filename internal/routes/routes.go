package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.POST("/", func(ctx *gin.Context) {
			log.Println("Creating task")
		})
	}

	return router
}