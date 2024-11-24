package dto

import "time"

type CreateTaskInput struct {
	Title string `json:"title"`
	Description *string `json:"description"`
	DueDate *time.Time `json:"due_date"`
}