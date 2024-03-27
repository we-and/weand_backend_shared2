package models

import (
	_ "time"

	"gorm.io/gorm"
)

type Ticketmessage struct {
	gorm.Model
	ID           uint32 `json:"id",gorm:"primaryKey"`
	UserId       uint32 `json:"user_id"`
	TicketId     uint32 `json:"ticket_id"`
	Content      string `json:"content"`
	IsRead       bool   `json:"is_read"`
	IsSentByUser bool   `json:"is_sent_by_user"`
}

func (c *Ticketmessage) GetId() uint32 {
	return c.ID
}

func (c *Ticketmessage) TableName() string {
	return "api_user.ticketmessage"
}
