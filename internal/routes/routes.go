package routes

import (
	"github.com/andrefelizardo/todo-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		// tasks.POST("/", controllers.CreateTask)
		// tasks.GET("/:id", controllers.FindTask)
		// tasks.PATCH("/:id", controllers.UpdateTask)
		// tasks.DELETE("/:id", controllers.DeleteTask)
	}

	userController := controllers.NewUserController()
	users := router.Group("/users")
	{
		users.POST("/", userController.CreateUser)
	}
	return router
}