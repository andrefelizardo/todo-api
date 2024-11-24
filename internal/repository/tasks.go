package repository

import (
	"database/sql"
	"time"

	"github.com/andrefelizardo/todo-api/internal/dto"
	"github.com/andrefelizardo/todo-api/internal/model"
	"github.com/google/uuid"
)

type TasksRepository interface {
	Create(task dto.CreateTaskInput) (model.Task, error)
}

type tasksRepositoryImpl struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) *tasksRepositoryImpl {
	return &tasksRepositoryImpl{db: db}
}

func (t *tasksRepositoryImpl) Create(task dto.CreateTaskInput) (model.Task, error) {
	id := uuid.New().String()
	var dueDate time.Time
	var now = time.Now()
	if task.DueDate == nil {
		dueDate = now
	} else {
		dueDate = *task.DueDate
	}
	_, err := t.db.Exec("INSERT into tasks (id, title, description, status, due_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
						id, task.Title, task.Description, model.StatusPending, dueDate, now, now)
	if err != nil {
		return model.Task{}, err
	}
	return model.Task{
		ID: id,
		Title: task.Title,
		Description: task.Description,
		DueDate: dueDate,
		Status: model.StatusPending,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}