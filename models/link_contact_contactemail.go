package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkContactContactemail struct {
	gorm.Model
	ID             int `json:"id",gorm:"primaryKey"`
	ContactemailId int `json:"contactemail_id`
	ContactId      int `json:"contact_id`
}

func (c *LinkContactContactemail) GetId() int {
	return c.ID
}

func (c *LinkContactContactemail) TableName() string {
	return "api_user.link_contact_contactemail"
}
