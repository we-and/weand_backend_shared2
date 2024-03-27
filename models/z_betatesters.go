package models

import (
	"gorm.io/gorm"
)

type Betatester struct {
	gorm.Model
	ID    uint32 `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (c *Betatester) GetId() uint32 {
	return c.ID
}

func (c *Betatester) TableName() string {
	return "api_monitoring.betatesters"
}
