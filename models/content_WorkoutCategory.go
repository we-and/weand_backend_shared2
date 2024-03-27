package models

import (
	"time"

	"gorm.io/gorm"
)


type WorkoutCategory struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Name string `json:"name"`
	IconUrl string `json:"icon_url"`
	Desc string `json:"desc"`
	Level uint32 `json:"level"`
	ParentId uint32 `json:"parent_id"`

	Children []WorkoutCategory `gorm:"foreignKey:parent_id" json:"links,omitempty"`
	//POPULATED
	LinksCategoryWorkout []LinkCategoryWorkout `gorm:"foreignKey:category_id" json:"links,omitempty"`
}

func (c *WorkoutCategory) GetId() uint32 {
	return c.ID
}

func (c *WorkoutCategory) TableName() string {
	return "api_content.workoutcategory"
}