package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type ContactAddress struct {
	gorm.Model
	ID         uint32     `json:"id",gorm:"primaryKey"`
	Address    string     `json:"address`
	Address2   string     `json:"address2`
	City       string     `json:"city`
	PostalCode string     `json:"postal_code`
	ContactId  uint32     `json:"contact_id`
	Country    string     `json:"country`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Deleted    bool       `json:"deleted,omitempty"`
}

func (c *ContactAddress) GetId() uint32 {
	return c.ID
}

func (c *ContactAddress) TableName() string {
	return "api_contact.contactaddress"
}
