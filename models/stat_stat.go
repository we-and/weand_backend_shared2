package models

import (
	"time"

	"gorm.io/gorm"
)

type Stat struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	gorm.Model
	Date_    time.Time      `json:"date"`
	Type string `json:"type"`
	Value   int `json:"value"`
	PublicValue  int `json:"public_value"`
}

func (c *Stat) GetId() uint32 {
	return c.ID
}

func (c *Stat) TableName() string {
	return "api_stats.global_stat"
}
