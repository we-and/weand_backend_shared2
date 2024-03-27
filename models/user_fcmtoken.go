package models

import (
	_ "time"

	"gorm.io/gorm"
)

type FirebaseFCMToken struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	UserId   uint32 `json:"user_id"`
	UserKey  string `json:"user_key"`
	PersonId   uint32 `json:"person_id"`
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}

func (c *FirebaseFCMToken) GetId() uint32 {
	return c.ID
}

func (c *FirebaseFCMToken) TableName() string {
	return "api_user.firebase_fcm_token"
}
