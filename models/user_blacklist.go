package models

import (
	"time"
	_ "time"
)

type RevokedJwt struct {
	//	gorm.Model
	ID            uint32     `json:"id",gorm:"primaryKey"`
	Token         string     `json:"token,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	Deleted       bool       `json:"deleted,omitempty"`
	RevokerUserId uint32     `json:"revoked_user_id,omitempty"`
	RevokeUserId  uint32     `json:"revoked_user_id,omitempty"`
}

func (c *RevokedJwt) GetId() uint32 {
	return c.ID
}

func (c *RevokedJwt) TableName() string {
	return "api_user.revoked_jwt"
}

func (c *RevokedJwt) SetDeleted() {
	c.Deleted = true
}
