package models

import (
	"time"

	"gorm.io/gorm"
)

type Audio struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Filename     string `json:"filename"`
	MoveId    uint32 `json:"move_id"`
	DescId    uint32 `json:"desc_id"`
	DescsetId    uint32 `json:"descset_id"`
	Voice    string `json:"voice"`
	Country     string `json:"country"`
	Type    string `json:"type"`
	DurationMs   int `json:"duration_ms"`
}

func (c *Audio) GetId() uint32 {
	return c.ID
}

func (c *Audio) TableName() string {
	return "api_content.audio"
}
