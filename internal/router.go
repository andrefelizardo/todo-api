package router

import (
	"database/sql"

	"github.com/andrefelizardo/todo-api/internal/handler"
	"github.com/andrefelizardo/todo-api/internal/repository"
	"github.com/andrefelizardo/todo-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	tasksRepo := repository.NewTasksRepository(db)
	tasksUsecase := usecase.NewTasksUsecase(tasksRepo)
	tasksHandler := handler.NewTasksHandler(*tasksUsecase)
	tasks := router.Group("/tasks")
	{
		tasks.POST("/", tasksHandler.Create)
		tasks.GET("/", tasksHandler.List)
		tasks.GET("/:id", tasksHandler.GetDetails)
	}

	return router
}