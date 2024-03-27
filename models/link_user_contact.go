package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkUserContact struct {
	gorm.Model
	ID        int `json:"id",gorm:"primaryKey"`
	UserId    int `json:"user_id`
	ContactId int `json:"contact_id`
}

func (c *LinkUserContact) GetId() int {
	return c.ID
}

func (c *LinkUserContact) TableName() string {
	return "api_contact.link_user_contact"
}
