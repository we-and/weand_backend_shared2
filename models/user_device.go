package models

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	ID      uint32 `json:"id"  gorm:"primaryKey"`
	UserID  uint32 `json:"user_id"`  //
	UserKey string `json:"user_key"` //
	//	Emodel       string     `json:"emodel"`
	Identifier   string     `json:"identifier"`
	Model_       string     `json:"model"`
	Manufacturer string     `json:"manufacturer"`
	Name         string     `json:"name"`
	Platform     string     `json:"platform"`
	Brand        string     `json:"brand"`
	OsName       string     `json:"os_"`
	OsVersion    string     `json:"os_version"`
	Extra        string     `json:"extra"`
	IsConfirmed  bool       `json:"is_confirmed"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	Deleted      bool       `json:"deleted,omitempty"`

	DeviceConfirmRequests []DeviceConfirmRequest `gorm:"foreignKey:device_id;references:identifier" json:"deviceconfirm,omitempty"`
}

func (c *Device) GetId() uint32 {
	return c.ID
}

func (c *Device) TableName() string {
	return "api_user.device"
}
