package models

import (
	"gorm.io/gorm"
)

type Onboarding struct {
	gorm.Model
	ID      uint32 `json:"id"  gorm:"primaryKey"`
	UserId  uint32 `json:"user_id"`  //
	Stage   string `json:"stage"`    //
	UserKey string `json:"user_key"` //
}

func (c *Onboarding) GetId() uint32 {
	return c.ID
}

func (c *Onboarding) TableName() string {
	return "api_progress.onboarding"
}
