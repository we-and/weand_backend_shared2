package models

import (
	"time"
)

type LinkUserOrg struct {
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Type      string     `json:"type"`
	UserId    uint32     `json:"user_id"`
	OrgId     uint32     `json:"org_id"`

	//POPULATED
	User *User `gorm:"foreignKey:user_id" json:"user,omitempty"`
	Org  *Team `gorm:"foreignKey:org_id" json:"org,omitempty"`
}

func (c *LinkUserOrg) GetId() uint32 {
	return c.ID
}

func (c *LinkUserOrg) TableName() string {
	return "api_team.link_user_org"
}

func (c *LinkUserOrg) isCoach() bool {
	return c.Type == "COACH"
}

func (c *LinkUserOrg) isMember() bool {
	return c.Type == "MEMBER"
}
