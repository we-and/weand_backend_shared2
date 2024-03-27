package models

import (
	"time"
)

type LinkEventPerson struct {
	ID               uint32     `json:"id"`
	EventId          uint32     `json:"event_id"`
	PersonId         uint32     `json:"person_id"`
	Reply            string     `json:"reply"`
	Shortmedium            string     `json:"shortmedium"`
	StartdatetimeUtc time.Time  `json:"startdatetime_utc"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`

	Person *Person `gorm:"foreignKey:person_id" json:"person,omitempty"`
	Event  *Event  `gorm:"foreignKey:event_id" json:"event,omitempty"`
}

func (c *LinkEventPerson) GetId() uint32 {
	return c.ID
}

func (c *LinkEventPerson) TableName() string {
	return "api_event.link_event_person"
}
