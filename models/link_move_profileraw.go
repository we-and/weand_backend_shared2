package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkMoveProfileraw struct {
	ID           uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt    *time.Time     `json:"created_at,omitempty"`
	UpdatedAt    *time.Time     `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"`
	MoveId       uint32         `json:"move_id"`
	ProfilerawId uint32         `json:"profileraw_id"`
	BatchId      string         `json:"batch_id"`
	SafetyLevel  string         `json:"safety_level"`
	Comment      string         `json:"comment"`

	
}

func (c *LinkMoveProfileraw) GetId() uint32 {
	return c.ID
}

func (c *LinkMoveProfileraw) TableName() string {
	return "api_content.link_move_profile"
}
