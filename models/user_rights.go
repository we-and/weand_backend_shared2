package models

import (
	_ "time"

	"gorm.io/gorm"
)

type Rights struct {
	gorm.Model
	ID     uint32 `json:"id" gorm:"primaryKey"`
	UserId uint32 `json:"user_id"`
	Rights string `json:"rights"`
}

func (c *Rights) GetId() uint32 {
	return c.ID
}

func (c *Rights) TableName() string {
	return "api_user.rights"
}
