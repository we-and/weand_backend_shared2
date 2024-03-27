package models

import (
	"gorm.io/gorm"
)

type LinkDeviceUser struct {
	gorm.Model
	ID       uint32 `json:"id"  gorm:"primaryKey"`
	UserId   uint32 `json:"user_id"`   //
	DeviceId uint32 `json:"device_id"` //
	UserKey  string `json:"user_key"`  //

	Device *Device `gorm:"foreignKey:device_id"`

}

func (c *LinkDeviceUser) GetId() uint32 {
	return c.ID
}

func (c *LinkDeviceUser) TableName() string {
	return "api_user.link_device_user"
}
