package model

import "time"

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	ID string
	Title string
	Description *string
	Status TaskStatus
	DueDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}