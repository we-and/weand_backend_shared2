package models

import (
	"time"

	"gorm.io/gorm"
)

type Physio struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx            uint32 `json:"idx"`
	Desc            string `json:"desc"`
	MoveId         uint32 `json:"move_id"`

	//POPULATED
	LinksAnatomy []LinkPhysioAnatomy `gorm:"foreignKey:physio_id" json:"anatomiczone,omitempty"`
	Move         *Move         `gorm:"foreignKey:move_id" json:"move,omitempty"`
}

func (c *Physio) GetId() uint32 {
	return c.ID
}

func (c *Physio) TableName() string {
	return "api_content.physio"
}
