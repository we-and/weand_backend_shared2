package models

import (
	"time"

	"gorm.io/gorm"
)

type Descset struct {
	ID         uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
	MoveId uint32 `json:"move_id"`
	Type string `json:"type"`


	LinkDescs []LinkDescsetDesc `gorm:"foreignKey:descset_id" json:"descset,omitempty"`
}

func (c *Descset) GetId() uint32 {
	return c.ID
}

func (c *Descset) TableName() string {
	return "api_content.descset"
}
