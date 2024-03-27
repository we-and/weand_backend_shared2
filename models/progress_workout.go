package models

import (
	"time"
)

type ProgressWorkout struct {
	ID        uint32 `json:"id"`
	WorkoutId uint32 `json:"workout_id"`

	StartedAt   time.Time  `json:"started_at,omitempty"`
	DurationSec int        `json:"duration_sec,omitempty"`
	EndedAt     time.Time  `json:"ended_at,omitempty"`
	UserKey     string     `json:"user_key,omitempty"`
	EndType     string     `json:"ended_type,omitempty"`
	InstanceId  string     `json:"instance_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	UserKeyObj UserKey `gorm:"foreignKey:user_key;references:user_key"`
}

func (c *ProgressWorkout) GetId() uint32 {
	return c.ID
}
func (c *ProgressWorkout) TableName() string {
	return "api_progress.workout"
}
