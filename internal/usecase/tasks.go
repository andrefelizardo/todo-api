package usecase

import (
	"github.com/andrefelizardo/todo-api/internal/dto"
	"github.com/andrefelizardo/todo-api/internal/model"
	"github.com/andrefelizardo/todo-api/internal/repository"
)

type TasksUsecase struct {
	repo repository.TasksRepository
}

func NewTasksUsecase(repo repository.TasksRepository) *TasksUsecase {
	return &TasksUsecase{repo: repo}
}

func (u *TasksUsecase) Create(input dto.CreateTaskInput) (model.Task, error) {
	task, err := u.repo.Create(input)
	if err != nil {
		return model.Task{}, err
	}
	
	return task, nil
}