package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkCategoryWorkout struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx    uint32 `json:"idx"`
	CategoryId    uint32 `json:"category_id"`
	WorkoutId uint32 `json:"workout_id"`

	//POPULATED
	Category    *Category  `gorm:"foreignKey:category_id" json:"category,omitempty"`
	Workout *Workout `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *LinkCategoryWorkout) GetId() uint32 {
	return c.ID
}

func (c *LinkCategoryWorkout) TableName() string {
	return "api_content.link_category_workout"
}
