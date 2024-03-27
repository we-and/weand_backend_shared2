package models

import (
	"time"

	"gorm.io/gorm"
)

type UsageStat struct {
	gorm.Model

	ID uint32 `json:"id" gorm:"primaryKey"`

	Date   time.Time `json:"date,omitempty"`
	TeamId uint32    `json:"team_id,omitempty"`
	Count  int    `json:"count,omitempty"`
	Usage  string    `json:"usage,omitempty"`
}

func (c *UsageStat) GetId() uint32 {
	return c.ID
}

func (c *UsageStat) TableName() string {
	return "api_job.usage"
}
