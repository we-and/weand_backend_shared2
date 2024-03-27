package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkoutProfile struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Conditions string `json:"conditions"`
	UserId     uint32 `json:"user_id"`
	UserKey    string `json:"user_key"`
	JobId      uint32 `json:"job_id"`
	Injuries   string `json:"injuries"`
	Gear       string `json:"gear"`
	Sports     string `json:"sports"`
	Goals      string `json:"goals"`
	Experience string `json:"experience"`
	Focus      string `json:"focus"`
}

func (c *WorkoutProfile) GetId() uint32 {
	return c.ID
}

func (c *WorkoutProfile) TableName() string {
	return "api_progress.workout_profile"
}
