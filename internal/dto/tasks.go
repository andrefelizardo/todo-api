package dto

import "time"

type CreateTaskInput struct {
	Title string `json:"title"`
	Description *string `json:"description"`
	// Status string `json:"status"`
	DueDate *time.Time `json:"due_date"`
}