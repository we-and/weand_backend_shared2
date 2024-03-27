package models

import (
	"gorm.io/gorm"
)

type UserSettings struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`
	//	CreatedAt *time.Time `json:"created_at,omitempty"`
	//	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserId     uint32 `gorm:"type:user_id" json:"user_id,omitempty"`
	CharacterKey    string        `json:"character_key,omitempty"`
	
}

func (c *UserSettings) GetId() uint32 {
	return c.ID
}

func (c *UserSettings) TableName() string {
	return "api_progress.user_settings"
}
