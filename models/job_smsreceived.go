package models

import (
	"time"

	"gorm.io/gorm"
)

type SmsReceived struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`

	Body             string     `json:"body,omitempty"`
	From         string     `json:"from,omitempty"`
	To       string     `json:"to,omitempty"`
	Isocode      string     `json:"isocode,omitempty"`
}

func (c *SmsReceived) GetId() uint32 {
	return c.ID
}

func (c *SmsReceived) TableName() string {
	return "api_job.sms_received"
}
