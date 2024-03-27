package models

import (
	"time"

	"gorm.io/gorm"
)

type SendJob struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Shortobject string     `json:"shortobject,omitempty"`
	Errdesc     string     `json:"errdesc,omitempty"`
	Senddate    *time.Time `json:"senddate,omitempty"`
	Started     bool       `json:"started,omitempty"`
	Sent        bool       `json:"sent,omitempty"`
	Success     bool       `json:"success,omitempty"`
	Shortmedium string     `json:"shortmedium,omitempty"`
	Destination string     `json:"destination,omitempty"`
	PersonId    uint32     `json:"person_id,omitempty"`
	TeamId      uint32     `json:"team_id,omitempty"`
	BatchId     uint32     `json:"batch_id,omitempty"`
	RelatedId   uint32     `json:"related_id,omitempty"`
}

func (c *SendJob) GetId() uint32 {
	return c.ID
}

func (c *SendJob) TableName() string {
	return "api_job.sendjob"
}
