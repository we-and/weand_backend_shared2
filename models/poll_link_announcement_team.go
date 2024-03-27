package models

import (
	"time"
)

type LinkAnnouncementTeam struct {
	ID             uint32 `json:"id"`
	AnnouncementId uint32 `json:"announcement_id"`
	TeamId         uint32 `json:"team_id"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Team         *Team         `gorm:"foreignKey:team_id" json:"team,omitempty"`
	Announcement *Announcement `gorm:"foreignKey:announcement_id" json:"announcement,omitempty"`
}

func (c *LinkAnnouncementTeam) GetId() uint32 {
	return c.ID
}

func (c *LinkAnnouncementTeam) TableName() string {
	return "api_poll.link_announcement_team"
}
