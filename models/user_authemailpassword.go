package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthEmailpassword struct {
	gorm.Model
	ID       uint32 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	//LastConfirmtoken string `json:"last_confirmtoken"`
	IsConfirmed bool       `json:"is_confirmed"`
	UserId      uint32     `json:"user_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Deleted     bool       `json:"deleted,omitempty"`

	EmailConfirmRequests []EmailConfirmRequest `gorm:"foreignKey:email;references:email" json:"emailconfirm,omitempty"`
}

func (c *AuthEmailpassword) GetId() uint32 {
	return c.ID
}

func (c *AuthEmailpassword) TableName() string {
	return "api_user.auth_emailpassword"
}
