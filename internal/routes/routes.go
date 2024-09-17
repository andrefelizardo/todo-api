package routes

import (
	"github.com/andrefelizardo/todo-api/internal/controllers"
	"github.com/andrefelizardo/todo-api/internal/middlewares"
	"github.com/andrefelizardo/todo-api/internal/repositories"
	"github.com/andrefelizardo/todo-api/internal/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//TODO: Tasks
	// [ ] Implement tasks' endpoints with authentication middleware
	// [ ] Implement task listing with user owner middleware
	//[ ] Implement task creation with verified middleware


	taskUsecase := usecases.NewTasksUseCase(repositories.NewTasksRepository(db))
	taskController := controllers.NewTasksController(*taskUsecase)

	tasks := router.Group("/tasks").Use(middlewares.AuthMiddleware())
	{
		tasks.GET("/", taskController.ListTasks)
	}
	// tasks := verified.Group("/tasks")
	// tasks.GET("/", taskController.ListTasks)
	
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

//TODO: Account creation v1
// [ ] Implement user creation
// [x] Generate JWT token on user creation and return
// [ ] Implement middleware to validate JWT token
// [ ] Implement middleware to validate if user is owner of resource
// [x] Implement middleware to validate if user has email verified
// [ ] Implement email confirmation on user creation

//TODO: Authentication flow v1
// [ ] Implement JWT authentication
// [ ] Implement user login
// [ ] Implement user logout



//TODO: User management v1
// [ ] Implement user password reset
// [ ] Implement user password change
// [ ] Implement user profile update
// [ ] Implement user profile delete
// [ ] Implement user profile view