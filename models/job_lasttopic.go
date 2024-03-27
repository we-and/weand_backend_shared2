package models

import (
	"time"

	"gorm.io/gorm"
)

type LastTopic struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	StartdatetimeUtc *time.Time `json:"startdatetime_utc"`
	RelatedId        uint32     `json:"related_id,omitempty"`
	TeamId           uint32     `json:"team_id,omitempty"`
	PersonId         uint32     `json:"person_id,omitempty"`
	RelatedObject    string     `json:"related_object,omitempty"`
	FromPhone        string     `json:"from_phone,omitempty"`
	WelcomeSent      bool       `json:"welcome_sent,omitempty"`
}

func (c *LastTopic) GetId() uint32 {
	return c.ID
}

func (c *LastTopic) TableName() string {
	return "api_job.last_topic"
}
