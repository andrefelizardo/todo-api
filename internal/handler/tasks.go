package handler

import (
	"net/http"

	"github.com/andrefelizardo/todo-api/internal/dto"
	"github.com/andrefelizardo/todo-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TasksHandler struct {
	tasksUsecase usecase.TasksUsecase
}

func NewTasksHandler(usecase usecase.TasksUsecase) *TasksHandler {
	return &TasksHandler{
		tasksUsecase: usecase,
	}
}

func (t *TasksHandler) Create(ctx *gin.Context) {
	var input dto.CreateTaskInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task input"})
		return
	}


	task, err := t.tasksUsecase.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": task})
}

func (t *TasksHandler) List(ctx *gin.Context) {
	tasks, err := t.tasksUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": tasks})
}