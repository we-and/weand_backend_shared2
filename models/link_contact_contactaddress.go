package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkContactContactaddress struct {
	gorm.Model
	ID               int `json:"id",gorm:"primaryKey"`
	ContactaddressId int `json:"contactaddress_id`
	ContactId        int `json:"contact_id`
}

func (c *LinkContactContactaddress) GetId() int {
	return c.ID
}

func (c *LinkContactContactaddress) TableName() string {
	return "api_contact.link_contact_contactaddress"
}
