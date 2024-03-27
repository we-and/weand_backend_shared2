package models

import (
	_ "time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	ID         uint32 `json:"id",gorm:"primaryKey"`
	UserId     uint32 `json:"user_id"`
	UserKey    uint32 `json:"user_key"`
	CategoryId uint32 `json:"category_id"`

	IsClosed      bool          `json:"is_closed"`
	LastMessageId int           `json:"last_message_id"`
	LastMessage   Ticketmessage `gorm:"foreignKey:last_message_id",json:"last_message,omitempty"`

	//POPULATED
	Ticketmessages []Ticketmessage `gorm:"foreignKey:ticket_id;references:id",json:"messages"`
}

func (c *Ticket) GetId() uint32 {
	return c.ID
}

func (c *Ticket) TableName() string {
	return "api_user.ticket"
}
