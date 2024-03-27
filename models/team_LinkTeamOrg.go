package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkTeamOrg struct {
	ID        uint32         `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	TeamId uint32 `json:"team_id"`
	OrgId  uint32 `json:"org_id"`

	//POPULATED
	Team *Team `gorm:"foreignKey:team_id" json:"team,omitempty"`
	Org  *Team `gorm:"foreignKey:org_id" json:"org,omitempty"`
}

func (c *LinkTeamOrg) GetId() uint32 {
	return c.ID
}

func (c *LinkTeamOrg) TableName() string {
	return "api_team.link_team_org"
}
