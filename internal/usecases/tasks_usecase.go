package usecases

import "github.com/andrefelizardo/todo-api/internal/repositories"

type TasksUseCase struct {
	tasksRepository repositories.TasksRepository
}

func NewTasksUseCase(tasksRepository repositories.TasksRepository) *TasksUseCase {
	return &TasksUseCase{
		tasksRepository: tasksRepository,
	}
}

func (t *TasksUseCase) List(userID string) ([]string, error) {
	return nil, nil
}