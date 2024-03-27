package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkMovegroupMove struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	BatchId   string         `json:"batch_id"`

	Idx         uint32 `json:"idx"`
	MovegroupId uint32 `json:"movegroup_id"`
	MoveId      uint32 `json:"move_id"`
	Side        string `json:"side"`
	Mode        string `json:"mode"`
	SafetyLevel string `json:"safety_level"`
	Comment     string `json:"comment"`

	//POPULATED
	Movegroup *Movegroup `gorm:"foreignKey:movegroup_id" json:"movegroup,omitempty"`
	Move      *Move      `gorm:"foreignKey:move_id" json:"move,omitempty"`
}

func (c *LinkMovegroupMove) GetId() uint32 {
	return c.ID
}

func (c *LinkMovegroupMove) TableName() string {
	return "api_content.link_movegroup_move"
}
