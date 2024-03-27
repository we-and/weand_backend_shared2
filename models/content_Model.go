package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Name     string `json:"name"`
	Url     string `json:"url"`
	AnimKey     string `json:"anim_key"`
	MoveId    uint32 `json:"move_id"`
	Filename     string `json:"filename"`
	Bucket     string `json:"bucket"`

	AxisUp      string `json:"axis_up,omitempty"`
	AssetId uint32 `json:"asset_id"`

	Asset *Asset `gorm:"foreignKey:asset_id" json:"asset,omitempty"`
	
}

func (c *Model) GetId() uint32 {
	return c.ID
}

func (c *Model) TableName() string {
	return "api_content.model"
}
