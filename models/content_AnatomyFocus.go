package models

import (
	"time"

	"gorm.io/gorm"
)

type Anatomyfocus struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	IsCustom  bool   `json:"is_custom"`
	AnatomyId uint32 `json:"anatomy_id"`
	Idx       uint32 `json:"idx"`
	Name      string `json:"name"`
}

func (c *Anatomyfocus) GetId() uint32 {
	return c.ID
}

func (c *Anatomyfocus) TableName() string {
	return "api_content.anatomyfocus"
}
