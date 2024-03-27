package models

import (
	"time"

	"gorm.io/gorm"
)

type Eventexception struct {
	gorm.Model
	ID        uint32     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	EventId   uint32     `json:"event_id"`

	StartdateUtc time.Time `json:"startdate_utc,omitempty"`

	//POPULATED FIELDS
	Event *Event `gorm:"foreignKey:event_id" json:"event,omitempty"`
}

func (c *Eventexception) GetId() uint32 {
	return c.ID
}

func (c *Eventexception) TableName() string {
	return "api_event.event_exception"
}
