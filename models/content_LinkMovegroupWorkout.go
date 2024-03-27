package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkMovegroupWorkout struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	BatchId string `json:"batch_id"`

	Idx         uint32 `json:"idx"`
	MovegroupId uint32 `json:"movegroup_id"`
	WorkoutId   uint32 `json:"workout_id"`

	//POPULATED
	Movegroup *Movegroup `gorm:"foreignKey:movegroup_id" json:"move,omitempty"`
	Workout   *Workout   `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *LinkMovegroupWorkout) GetId() uint32 {
	return c.ID
}

func (c *LinkMovegroupWorkout) TableName() string {
	return "api_content.link_movegroup_workout"
}
