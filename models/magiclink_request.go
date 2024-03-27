package models

import (
	"time"

	"gorm.io/gorm"
)

type MagiclinkRequest struct {
	gorm.Model
	ID             uint32     `json:"id"`
	UserId         uint32     `json:"user_id"`
	Email          string     `json:"email"`
	MagiclinkToken string     `json:"magiclink_token"`
	PublicToken    string     `json:"public_token"`
	Link           string     `json:"link"`
	Sent           bool       `json:"sent`
	Confirmed      bool       `json:"confirmed`
	Renewed        bool       `json:"renewed`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}

func (c *MagiclinkRequest) GetId() uint32 {
	return c.ID
}

func (c *MagiclinkRequest) TableName() string {
	return "api_user.request_magiclink"
}
