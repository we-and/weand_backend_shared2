package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkSocialConfirmRequest struct {
	gorm.Model
	ID        uint32     `json:"id"  gorm:"primaryKey"`
	Email     string     `json:"email"`
	Token     string     `json:"token"`
	Link      string     `json:"link"`
	SocialId  int        `json:"social_id"`
	Trigger   string     `json:"trigger"`
	Sent      bool       `json:"sent"`
	Confirmed bool       `json:"confirmed"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *LinkSocialConfirmRequest) GetId() uint32 {
	return c.ID
}

func (c *LinkSocialConfirmRequest) TableName() string {
	return "api_request.request_linksocialconfirm"
}
