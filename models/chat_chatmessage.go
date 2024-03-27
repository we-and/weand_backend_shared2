package models

import (
	_ "time"

	"gorm.io/gorm"
)

type Chatmessage struct {
	gorm.Model
	ID              uint32 `json:"id" gorm:"primaryKey"`
	CreatorPersonId uint32 `json:"creator_person_id"`
	ThreadId        uint32 `json:"thread_id"`
	TeamId          uint32 `json:"team_id"`
	Content         string `json:"content"`
	Type            string `json:"type" gorm:"type:char"`
	IsSoftDeleted   bool   `json:"is_soft_deleted"`

	//POPULATED
	Thread  *Thread                  `gorm:"foreignKey:thread_id;references:id" json:"thread"`
	Creator *Person                  `gorm:"foreignKey:creator_person_id;references:id" json:"creator"`
	Viewers *[]LinkChatmessageViewer `gorm:"foreignKey:message_id;references:id" json:"viewers"`
}

func (c *Chatmessage) GetId() uint32 {
	return c.ID
}

func (c *Chatmessage) TableName() string {
	return "api_chat.chat_message"
}
