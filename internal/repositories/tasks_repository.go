package repositories

import (
	"github.com/andrefelizardo/todo-api/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TasksRepository interface {
	List(userID uuid.UUID) ([]domain.Task, error)
}

type tasksRepositoryImpl struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) *tasksRepositoryImpl {
	return &tasksRepositoryImpl{
		db: db,
	}
}

func (t *tasksRepositoryImpl) List(userID uuid.UUID) ([]domain.Task, error) {
	var tasks []domain.Task
	result := t.db.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
