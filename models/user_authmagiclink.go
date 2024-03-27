package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthMagiclink struct {
	gorm.Model
	ID        uint32     `json:"id"`
	Email     string     `json:"email"`
	Key       string     `json:"key"`
	UserId    uint32     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//User User `gorm:"foreignKey:user_id"`
}

func (c *AuthMagiclink) GetId() uint32 {
	return c.ID
}

func (c *AuthMagiclink) TableName() string {
	return "api_user.auth_magiclink"
}
