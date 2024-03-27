package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Trainingtimes struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`
	UserId     uint32 `json:"user_id,omitempty"`
	UserKey     string `json:"user_key,omitempty"`
	Recap    pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null" json:"recap,omitempty"`
	
}

func (c *Trainingtimes) GetId() uint32 {
	return c.ID
}

func (c *Trainingtimes) TableName() string {
	return "api_progress.trainingtimes"
}
