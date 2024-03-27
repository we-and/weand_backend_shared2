package models

import (
	"time"
)

type LinkEventTeam struct {
	ID      uint32 `json:"id"`
	EventId uint32 `json:"event_id"`
	TeamId  uint32 `json:"team_id"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Team  *Team  `gorm:"foreignKey:team_id" json:"team,omitempty"`
	Event *Event `gorm:"foreignKey:event_id" json:"event,omitempty"`
}

func (c *LinkEventTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkEventTeam) TableName() string {
	return "api_event.link_event_team"
}
