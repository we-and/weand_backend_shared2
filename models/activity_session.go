package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	UserId   uint32 `json:"user_id"`
	DeviceId uint32 `json:"device_id"`

	Success   bool       `json:"success"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *Session) GetId() uint32 {
	return c.ID
}

func (c *Session) TableName() string {
	return "api_feed.session"
}
