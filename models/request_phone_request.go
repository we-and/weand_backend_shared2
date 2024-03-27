package models

import (
	"time"

	"gorm.io/gorm"
)

type PhoneRequest struct {
	gorm.Model
	ID          uint32     `json:"id" gorm:"primaryKey"`	
	UserId      uint32     `json:"user_id"`
	PhoneId     uint32     `json:"phone_id"`
	Phone       string     `json:"phone"`
	CountryCode string     `json:"country_code"`
	Sent        bool       `json:"sent"`
	Code        string     `json:"code"`
	Confirmed   bool       `json:"confirmed"`
	Renewed     bool       `json:"renewed"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	User User `gorm:"foreignKey:user_id"`
}

func (c *PhoneRequest) GetId() uint32 {
	return c.ID
}

func (c *PhoneRequest) TableName() string {
	return "api_request.request_phone"
}
