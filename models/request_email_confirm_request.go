package models

import (
	"time"

	"gorm.io/gorm"
)

type EmailConfirmRequest struct {
	gorm.Model
	ID          uint32        `json:"id" gorm:"primaryKey"`
	Email       string     `json:"email"`
	Token       string     `json:"token"`
	Trigger     string     `json:"trigger"`
	Link        string     `json:"link"`
	AppLink     string     `json:"applink"`
	Sent        bool       `json:"sent"`
	Confirmed   bool       `json:"confirmed"`
	PasswordSet bool       `json:"passwordset"`
	Renewed     bool       `json:"renewed"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (c *EmailConfirmRequest) GetId() uint32 {
	return c.ID
}

func (c *EmailConfirmRequest) TableName() string {
	return "api_request.request_emailconfirm"
}
