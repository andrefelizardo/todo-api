package model

import "time"

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted TaskStatus = "completed"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description *string `json:"description"`
	Status TaskStatus `json:"status"`
	DueDate time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}