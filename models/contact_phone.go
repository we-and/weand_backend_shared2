package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type ContactPhone struct {
	gorm.Model
	ID        uint32     `json:"id",gorm:"primaryKey"`
	Phone     string     `json:"phone`
	ContactId uint32     `json:"contact_id`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Deleted   bool       `json:"deleted,omitempty"`
}

func (c *ContactPhone) GetId() uint32 {
	return c.ID
}

func (c *ContactPhone) TableName() string {
	return "api_contact.contactphone"
}
