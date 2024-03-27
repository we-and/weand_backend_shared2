package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkTeamCategory struct {
	gorm.Model
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	TeamId     uint32 `json:"team_id"`
	CategoryId uint32 `json:"category_id"`

	//POPULATED
	Team     *Team     `gorm:"foreignKey:team_id",json:"team,omitempty"`
	Category *Category `gorm:"foreignKey:category_id",json:"category,omitempty"`
}

func (c *LinkTeamCategory) GetId() uint32 {
	return c.ID
}

func (c *LinkTeamCategory) TableName() string {
	return "api_team.link_team_category"
}
