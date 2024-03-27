package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthPhone struct {
	gorm.Model
	ID          uint32 `json:"id"`
	UserId      uint32 `json:"user_id"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
	Confirmed   bool   `json:"confirmed`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	User      *User      `gorm:"foreignKey:user_id"`
}

func (c *AuthPhone) GetId() uint32 {
	return c.ID
}

func (c *AuthPhone) TableName() string {
	return "api_user.auth_phone"
}
