package controllers

import (
	"net/http"

	"github.com/andrefelizardo/todo-api/internal/middlewares"
	"github.com/andrefelizardo/todo-api/internal/usecases"
	"github.com/gin-gonic/gin"
)

type TasksController struct {
	tasksUsecase usecases.TasksUseCase
}

func NewTasksController(tasksUsecase usecases.TasksUseCase) *TasksController {
	return &TasksController{
		tasksUsecase: tasksUsecase,
	}
}

func (t *TasksController) ListTasks(ctx *gin.Context) {
	userID, exists := ctx.Get(middlewares.UserContextKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	tasks, err := t.tasksUsecase.List(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tasks})

}