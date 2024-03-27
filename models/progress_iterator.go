package models

import (
	"time"
)

type Iterator struct {
	ID          uint32     `json:"id"`
	UserKey     string     `json:"user_key,omitempty"`
	UserId      uint32     `json:"user_id,omitempty"`
	ProgramId   uint32     `json:"program_id,omitempty"`
	WorkoutId   uint32     `json:"workout_id,omitempty"`
	PlangroupId uint32     `json:"plangroup_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	Program *Program `gorm:"foreignKey:program_id" json:"program,omitempty"`
}

func (c *Iterator) GetId() uint32 {
	return c.ID
}
func (c *Iterator) TableName() string {
	return "api_progress.iterator"
}
