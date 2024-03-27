package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Points   int               `json:"points"`
	PointsEarned   int               `json:"points_earned"`
	Idx         uint32     `json:"idx"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	LessonId    uint32     `json:"lesson_id"`
	Type        string     `json:"type"`
	ItemId      uint32     `json:"item_id"`
	WorkoutId      uint32     `json:"workout_id"`
	ReadingId      uint32     `json:"reading_id"`
	Title       string     `json:"title"`
	ItemWorkout *Workout   `gorm:"foreignKey:workout_id" `
	ItemReading *Article   `gorm:"foreignKey:reading_id" `
}

func (c *Task) GetId() uint32 {
	return c.ID
}
func (c *Task) TableName() string {
	return "api_progress.task2"
}
