package routes

import (
	"github.com/andrefelizardo/todo-api/internal/controllers"
	"github.com/andrefelizardo/todo-api/internal/repositories"
	"github.com/andrefelizardo/todo-api/internal/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// tasks := router.Group("/tasks")
	// {
	// 	tasks.GET("/", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "pong",
	// 		})
	// 	})
		// tasks.POST("/", controllers.CreateTask)
		// tasks.GET("/:id", controllers.FindTask)
		// tasks.PATCH("/:id", controllers.UpdateTask)
		// tasks.DELETE("/:id", controllers.DeleteTask)
	// }

	userUsecase := usecases.NewUserUseCase(repositories.NewUserRepository(db))
	userController := controllers.NewUserController(*userUsecase)
	users := router.Group("/users")
	{
		users.POST("/", userController.CreateUser)
	}
	return router
}