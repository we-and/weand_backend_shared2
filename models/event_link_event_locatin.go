package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkEventLocation struct {
	gorm.Model
	ID             uint32 `json:"id"`
	EventId        uint32 `json:"event_id"`
	TeamlocationId uint32 `json:"teamlocation_id"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Location *TeamLocation `gorm:"foreignKey:TeamlocationId;references:ID" json:"location,omitempty"`
	Event    *Event        `gorm:"foreignKey:EventId;references:ID" json:"event,omitempty"`
}

func (c *LinkEventLocation) GetId() uint32 {
	return c.ID
}

func (c *LinkEventLocation) TableName() string {
	return "api_event.link_event_location"
}
