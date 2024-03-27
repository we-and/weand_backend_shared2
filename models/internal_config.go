package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	Value     string     `json:"value"`
	Key string     `json:"key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
func (c *Config) GetId() uint32 {
	return c.ID
}

func (c *Config) TableName() string {
	return "api_internal.config"
}
