package models

import (
	"time"

	"gorm.io/gorm"
)

type Translation struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Lang   string `json:"lang"`
	MoveId uint32 `json:"move_id"`
	ObjId  uint32 `json:"obj_id"`
	Type   string `json:"type"`
	Text   string `json:"text"`
}

func (c *Translation) GetId() uint32 {
	return c.ID
}

func (c *Translation) TableName() string {
	return "api_content.translation"
}
