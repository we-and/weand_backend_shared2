package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkUserWorkoutprofile struct {
	ID           uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt    *time.Time     `json:"created_at,omitempty"`
	UpdatedAt    *time.Time     `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
	UserId       uint32         `json:"user_id"`
	UserKey      string         `json:"user_key"`
	ProfilerawId uint32         `json:"profileraw_id"`

	WorkoutProfileRaw *WorkoutProfileRaw `gorm:"foreignKey:profileraw_id;references:id" json:"workoutprofileraw,omitempty"`
}

func (c *LinkUserWorkoutprofile) GetId() uint32 {
	return c.ID
}

func (c *LinkUserWorkoutprofile) TableName() string {
	return "api_survey.link_user_profile"
}
