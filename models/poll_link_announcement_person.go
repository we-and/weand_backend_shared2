package models

import (
	"time"
)

type LinkAnnouncementPerson struct {
	ID             uint32 `json:"id"`
	AnnouncementId uint32 `json:"announcement_id"`
	PersonId       uint32 `json:"person_id"`
	Reply          string `json:"reply"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Person       *Person       `gorm:"foreignKey:person_id" json:"person,omitempty"`
	Announcement *Announcement `gorm:"foreignKey:announcement_id" json:"announcement,omitempty"`
}

func (c *LinkAnnouncementPerson) GetId() uint32 {
	return c.ID
}
func (c *LinkAnnouncementPerson) TableName() string {
	return "api_poll.link_announcement_user"
}
