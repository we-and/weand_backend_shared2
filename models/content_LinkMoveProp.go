package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkMoveProp struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	//Idx     uint32 `json:"idx"`
	AssetId uint32 `json:"asset_id"`
	MoveId  uint32 `json:"move_id"`
	IsMain  bool   `json:"is_main"`

	//POPULATED
	Asset *Asset `gorm:"foreignKey:asset_id" json:"asset,omitempty"`
	Move  *Move  `gorm:"foreignKey:move_id" json:"move,omitempty"`
}

func (c *LinkMoveProp) GetId() uint32 {
	return c.ID
}

func (c *LinkMoveProp) TableName() string {
	return "api_content.link_move_prop"
}
