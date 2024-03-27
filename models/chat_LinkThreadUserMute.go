package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkThreadUserMute struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	ThreadId uint32 `json:"thread_id"`
	UserId   uint32 `json:"user_id"`
	Duration uint32 `json:"duration"`

	//POPULATED FIELDS
	User *User `gorm:"foreignKey:user_id" json:"user,omitempty"`
}

func (c *LinkThreadUserMute) GetId() uint32 {
	return c.ID
}

func (c *LinkThreadUserMute) TableName() string {
	return "api_chat.link_thread_user_mute"
}
