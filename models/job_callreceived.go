package models

import (
	"time"

	"gorm.io/gorm"
)

type CallReceived struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`

	Body             string     `json:"body,omitempty"`
	From         string     `json:"from,omitempty"`
	To       string     `json:"to,omitempty"`
	Isocode      string     `json:"isocode,omitempty"`
}

func (c *CallReceived) GetId() uint32 {
	return c.ID
}

func (c *CallReceived) TableName() string {
	return "api_job.call_received"
}
