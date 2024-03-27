package models

import (
	"time"

	"gorm.io/gorm"
)

type DeviceConfirmRequest struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	Token     string     `json:"token"`
	DeviceId  string     `json:"device_id"`
	Trigger   string     `json:"trigger"`
	Link      string     `json:"link"`
	Sent      bool       `json:"sent"`
	Confirmed bool       `json:"confirmed"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *DeviceConfirmRequest) GetId() uint32 {
	return c.ID
}

func (c *DeviceConfirmRequest) TableName() string {
	return "api_request.request_deviceconfirm"
}
