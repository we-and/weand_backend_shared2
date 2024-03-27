package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthStrategy struct {
	gorm.Model
	ID                  uint32 `json:"id"`
	UserId              uint32 `json:"user_id"`
	AuthFacebookId      uint32 `json:"auth_facebook_id"`
	AuthGoogleId        uint32 `json:"auth_google_id"`
	AuthEmailpasswordId uint32 `json:"auth_emailpassword_id"`
	AuthAppleId         uint32 `json:"auth_apple_id"`
	//Strategy            string `json:"strategy"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *AuthStrategy) GetId() uint32 {
	return c.ID
}

func (c *AuthStrategy) TableName() string {
	return "api_user.auth_strategy"
}
