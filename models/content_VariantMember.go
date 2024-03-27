package models

import (
	"gorm.io/gorm"
)

type VariantMember struct {
	gorm.Model
	ID            uint32 `json:"id"  gorm:"primaryKey"`
	GroupId uint32 `json:"group_id"` //
	MoveId uint32 `json:"move_id"` //
	
	RelatedMove Move `gorm:"foreignKey:move_id;references:id" json:"move,omitempty"`
	Group *VariantGroup `gorm:"foreignKey:group_id" json:"group,omitempty"`
}

func (c *VariantMember) GetId() uint32 {
	return c.ID
}

func (c *VariantMember) TableName() string {
	return "api_content.variantmember"
}
