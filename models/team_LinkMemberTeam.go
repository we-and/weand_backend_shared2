package models

import (
	"time"
)

type LinkMemberTeam struct {
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserId uint32 `json:"user_id"`
	TeamId uint32 `json:"team_id"`

	//POPULATED
	Member *User `gorm:"foreignKey:user_id",json:"member,omitempty"`
	Team   *Team `gorm:"foreignKey:team_id",json:"team,omitempty"`
}

func (c *LinkMemberTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkMemberTeam) TableName() string {
	return "api_team.link_member_team"
}
