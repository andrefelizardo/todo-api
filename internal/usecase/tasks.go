package usecase

import (
	"errors"

	"github.com/andrefelizardo/todo-api/internal/dto"
	"github.com/andrefelizardo/todo-api/internal/model"
	"github.com/andrefelizardo/todo-api/internal/repository"
	"github.com/google/uuid"
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

func (u *TasksUsecase) FindAll() ([]model.Task, error) {
	tasks, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u *TasksUsecase) FindByID(id string) (model.Task, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return model.Task{}, errors.New("invalid ID")
	}

	task, err := u.repo.FindByID(id)
	if err != nil {
		return model.Task{}, err
	}

	return task, nil
}