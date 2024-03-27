package models

import (
	"time"
)

type LinkNicknamePersonTeam struct {
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	CreatorPersonId   uint32 `json:"creator_person_id"`
	TeamId            uint32 `json:"team_id"`
	NicknamedPersonId uint32 `json:"nicknamed_person_id"`
	Nickname          string `json:"nickname"`

	//POPULATED
	Person    *User `gorm:"foreignKey:person_id" json:"person,omitempty"`
	Nicknamed *User `gorm:"foreignKey:nicknamed_id" json:"nicknamed,omitempty"`
	Team      *Team `gorm:"foreignKey:team_id" json:"team,omitempty"`
}

func (c *LinkNicknamePersonTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkNicknamePersonTeam) TableName() string {
	return "api_team.link_nickname_user_team"
}
