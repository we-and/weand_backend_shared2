package models

import (
	"gorm.io/gorm"
)

type VariantGroup struct {
	gorm.Model
	ID           uint32          `json:"id"  gorm:"primaryKey"`
	MasterMoveId uint32          `json:"master_move_id"` //
	Key          string          `json:"key"`            //
	GroupMembers []VariantMember `gorm:"foreignKey:group_id" json:"members,omitempty"`
}

func (c *VariantGroup) GetId() uint32 {
	return c.ID
}

func (c *VariantGroup) TableName() string {
	return "api_content.variantgroup"
}
