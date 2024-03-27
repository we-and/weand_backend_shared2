package models

import (
	"time"
)

type LinkPollTeam struct {
	ID     uint32 `json:"id"`
	PollId uint32 `json:"poll_id"`
	TeamId uint32 `json:"team_id"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Team *Team `gorm:"foreignKey:team_id" json:"team,omitempty"`
	Poll *Poll `gorm:"foreignKey:poll_id" json:"Poll,omitempty"`
}

func (c *LinkPollTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkPollTeam) TableName() string {
	return "api_poll.link_poll_team"
}
