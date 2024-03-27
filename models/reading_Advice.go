package models

import (
	"time"

	"gorm.io/gorm"
)

type Advice struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Title         string `json:"title"`
	Content       string `json:"content"`
	DurationMin  int    `json:"duration_min"`

}

func (c *Advice) GetId() uint32 {
	return c.ID
}

func (c *Advice) TableName() string {
	return "api_reading.advice"
}
