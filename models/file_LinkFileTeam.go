package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type LinkFileTeam struct {
	gorm.Model
	ID        uint32 `json:"id" gorm:"primaryKey"`
	TeamID    uint32 `json:"team_id" gorm:"primaryKey"`
	ProgramID uint32 `json:"file_id" gorm:"primaryKey"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	File *File `gorm:"foreignKey:file_id" json:"file,omitempty"`
	Team *Team `gorm:"foreignKey:team_id" json:"team,omitempty"`
}

func (c *LinkFileTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkFileTeam) TableName() string {
	return "api_file.link_file_team"
}
