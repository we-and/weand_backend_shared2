package models

import (
	"stretches-common-api/structs"
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	ID        uint32     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Title   string `json:"title"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Content string `json:"content"`

	//POPULATED FIELDS
	LinkTeams []LinkAnnouncementTeam   `gorm:"foreignKey:announcement_id" json:"linkteams,omitempty"`
	Viewers   []LinkAnnouncementPerson `gorm:"foreignKey:announcement_id" json:"viewers,omitempty"`
	Batches   []SendBatch              `gorm:"foreignKey:related_id" json:"batches,omitempty"`
}

func (c *Announcement) GetId() uint32 {
	return c.ID
}

func (c *Announcement) TableName() string {
	return "api_poll.announcement"
}

func (c *Announcement) CanEdit(me structs.Me) bool {
	for _, liTeam := range c.LinkTeams {
		if liTeam.Team != nil {
			if liTeam.Team.CanEdit(me) {
				return true
			}
		}
	}
	return false
}
