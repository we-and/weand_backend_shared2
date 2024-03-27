package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type ContactEmail struct {
	gorm.Model
	ID        uint32     `json:"id",gorm:"primaryKey"`
	Email     string     `json:"email`
	ContactId uint32     `json:"contact_id`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Deleted   bool       `json:"deleted,omitempty"`
}

func (c *ContactEmail) GetId() uint32 {
	return c.ID
}

func (c *ContactEmail) TableName() string {
	return "api_contact.contactemail"
}
