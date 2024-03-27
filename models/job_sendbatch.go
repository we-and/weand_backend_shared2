package models

import (
	"time"

	"gorm.io/gorm"
)

type SendBatch struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Shortobject             string     `json:"shortobject,omitempty"`
	TeamId                  uint32     `json:"team_id,omitempty"`
	PersonId                uint32     `json:"person_id,omitempty"`
	RelatedId               uint32     `json:"related_id,omitempty"`
	RelatedStartdatetimeUtc *time.Time `json:"related_startdatetime_utc,omitempty"`
	Jobs                    []SendJob  `gorm:"foreignKey:batch_id" json:"jobs,omitempty"`
}

func (c *SendBatch) GetId() uint32 {
	return c.ID
}

func (c *SendBatch) TableName() string {
	return "api_job.sendbatch"
}
