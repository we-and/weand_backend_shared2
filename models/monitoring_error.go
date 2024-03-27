package models

import (
	"time"
)

type MonitoringError struct {
	ID      int        `json:"id",gorm:"primaryKey"`
	When    *time.Time `json:"when,omitempty"`
	Route   string     `json:"route,omitempty"`
	Type    string     `json:"type,omitempty"`
	Trigger string     `json:"trigger,omitempty"`
	Desc    string     `json:"desc,omitempty"`
	Backend string     `json:"backend,omitempty"`
	Code    string     `json:"code,omitempty"`
	UserId  *int       `json:"user_id,omitempty"`
	User    User       `gorm:"foreignKey:user_id"`
}

func (c *MonitoringError) GetId() int {
	return c.ID
}

func (c *MonitoringError) TableName() string {
	return "api_monitoring.errors"
}
