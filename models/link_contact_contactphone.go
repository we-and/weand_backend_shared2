package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkContactContactphone struct {
	gorm.Model
	ID             int `json:"id",gorm:"primaryKey"`
	ContactphoneId int `json:"contactphone_id`
	ContactId      int `json:"contact_id`
}

func (c *LinkContactContactphone) GetId() int {
	return c.ID
}

func (c *LinkContactContactphone) TableName() string {
	return "api_user.link_contact_contactphone"
}
