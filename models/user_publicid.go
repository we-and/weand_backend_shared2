package models

import (
	_ "time"

	"gorm.io/gorm"
)

type PublicId struct {
	gorm.Model
	ID        uint32 `json:"id"`
	UserID    uint32 `json:"user_id" gorm:"primaryKey"`
	Publickey string `json:"publickey"`
}

func (c *PublicId) GetId() uint32 {
	return c.ID
}

func (c *PublicId) TableName() string {
	return "api_user.publicid"
}
