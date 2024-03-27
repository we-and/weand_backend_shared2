package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkTagWorkout struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx    uint32 `json:"idx"`
	TagId    uint32 `json:"tag_id"`
	WorkoutId uint32 `json:"workout_id"`

	//POPULATED
	Tag   *Tag   `gorm:"foreignKey:tag_id" json:"tag,omitempty"`
	Workout *Workout `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *LinkTagWorkout) GetId() uint32 {
	return c.ID
}

func (c *LinkTagWorkout) TableName() string {
	return "api_content.link_tag_workout"
}
