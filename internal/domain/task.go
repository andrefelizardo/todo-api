package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key"`
	Title string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Status string `gorm:"type:varchar(20);not null;check:status IN ('pending', 'in_progress', 'completed')"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index;constraint:OnDelete:CASCADE"`
	DueDate time.Time `gorm:"type:date"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}