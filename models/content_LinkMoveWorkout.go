package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkMoveWorkout struct {
	ID        uint32         `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx    uint32 `json:"idx"`
	MoveId    uint32 `json:"move_id"`
	WorkoutId uint32 `json:"workout_id"`

	//POPULATED
	Move    *Move   `gorm:"foreignKey:move_id" json:"move,omitempty"`
	Workout *Workout `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *LinkMoveWorkout) GetId() uint32 {
	return c.ID
}

func (c *LinkMoveWorkout) TableName() string {
	return "api_content.link_move_workout"
}
