package models

import (
	"time"
)

type Challenge struct {
	ID                   uint32                   `json:"id"`
	Name               string                   `json:"user_key,omitempty"`
	IsActive bool  `json:"is_active"`
	CreatedAt            *time.Time               `json:"created_at,omitempty"`
	StartDate            time.Time               `json:"start_date,omitempty"`
	EndDate            time.Time               `json:"end_date,omitempty"`
	UpdatedAt            *time.Time               `json:"updated_at,omitempty"`
	DeletedAt            *time.Time               `json:"deleted_at,omitempty"`
}

func (c *Challenge) GetId() uint32 {
	return c.ID
}
func (c *Challenge) TableName() string {
	return "api_progress.challenge"
}
