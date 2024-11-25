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
	FindAll() ([]model.Task, error)
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

func (t *tasksRepositoryImpl) FindAll() ([]model.Task, error) {
	rows, err := t.db.Query("SELECT id, title, description, status, due_date, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var (
			id, title, status, due_date, created_at, updated_at string
			description sql.NullString
		)
		if err := rows.Scan(&id, &title, &description, &status, &due_date, &created_at, &updated_at); err != nil {
			return nil, err
		}


		dueDate, err := dateConvert(due_date)
		if err != nil {
			
		}

		createdAt, err := dateConvert(created_at)
		if err != nil {
			return nil, err
		}

		updatedAt, err := dateConvert(updated_at)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, model.Task{
			ID: id,
			Title: title,
			Description: &description.String,
			Status: model.TaskStatus(status),
			DueDate: dueDate,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return tasks, nil
}

func dateConvert(input string) (time.Time, error) {
	dateLayout := "2006-01-02 15:04:05.999999-07:00"
	date, err := time.Parse(dateLayout, input)
		if err != nil {
			return time.Time{}, err
		}
		return date, nil
}