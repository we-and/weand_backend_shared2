package models

import (
	"time"
)

type AppEvent struct {
	ID        uint32     `json:"id"`
	Type      string     `json:"ended_type,omitempty"`
	ExtraId   int        `json:"extra_id,omitempty"`
	UserKey   string     `json:"user_key,omitempty"`
	UserId    uint32     `json:"user_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *AppEvent) GetId() uint32 {
	return c.ID
}
func (c *AppEvent) TableName() string {
	return "api_progress.app_event"
}
