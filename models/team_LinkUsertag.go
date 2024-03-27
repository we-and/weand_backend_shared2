package models

import (
	"gorm.io/gorm"
)

type LinkPersonTag struct {
	gorm.Model
	ID              uint32 `json:"id" gorm:"primaryKey"`
	PersonId        uint32 `json:"person_id"`
	TeamId          uint32 `json:"team_id"`
	TagId           uint32 `json:"tag_id"`
	CreatorPersonId uint32 `json:"creator_person_id"`

	//POPULATED
	Tag  *Persontag `gorm:"foreignKey:tag_id" json:"tag,omitempty"`
	Team *Team      `gorm:"foreignKey:team_id" json:"team,omitempty"`
}

func (c *LinkPersonTag) GetId() uint32 {
	return c.ID
}

func (c *LinkPersonTag) TableName() string {
	return "api_team.link_person_tag"
}
