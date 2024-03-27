package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkChatmessageViewer struct {
	gorm.Model
	ID        uint32 `json:"id" gorm:"primaryKey"`
	MessageId uint32 `json:"message_id"`
	PersonId  uint32 `json:"person_id"`

	//POPULATED FIELDS
	Person *Person `gorm:"foreignKey:user_id" json:"person,omitempty"`
}

func (c *LinkChatmessageViewer) GetId() uint32 {
	return c.ID
}

func (c *LinkChatmessageViewer) TableName() string {
	return "api_chat.link_message_user_viewers"
}
