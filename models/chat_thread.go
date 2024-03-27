package models

import (
	_ "time"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	ID                 uint32 `json:"id" gorm:"primaryKey"`
	CreatorPersonId    uint32 `json:"creator_person_id"`
	TeamId             uint32 `json:"team_id"`
	IsArchived         bool   `json:"is_archived"`
	Name               string `json:"name"`
	UserKey            string `json:"user_key"`
	Hashtag            string `json:"hashtag"`
	Type               string `json:"type"` ///all ,
	LastMessageId      uint32 `json:"last_message_id"`
	LastMessageExtract string `json:"last_message_extract"`
	UnreadCount        int    `json:"unread_count"`
	//POPULATED
	LastMessage   *Chatmessage       `gorm:"foreignKey:last_message_id" json:"last_message,omitempty"`
	Messages      []Chatmessage      `gorm:"foreignKey:thread_id;references:id" json:"messages"`
	CustomMembers []LinkThreadMember `gorm:"foreignKey:thread_id;references:id" json:"members_custom"`
}

func (c *Thread) GetId() uint32 {
	return c.ID
}

func (c *Thread) TableName() string {
	return "api_chat.thread"
}
