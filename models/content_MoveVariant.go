package models

import (
	"gorm.io/gorm"
)

type MoveVariant struct {
	gorm.Model
	ID            uint32 `json:"id"  gorm:"primaryKey"`
	MoveId        uint32 `json:"move_id"`         //
	VariantMoveId uint32 `json:"variant_move_id"` //
	Level         string `json:"level"`           //
	Name          string `json:"name"`            //
	Desc          string `json:"desc"`            //

	MoveVariant *Move `gorm:"foreignKey:variant_move_id" json:"move,omitempty"`
}

func (c *MoveVariant) GetId() uint32 {
	return c.ID
}

func (c *MoveVariant) TableName() string {
	return "api_content.movevariant"
}
