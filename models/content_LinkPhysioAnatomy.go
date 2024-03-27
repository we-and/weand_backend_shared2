package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkPhysioAnatomy struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	PhysioId  uint32 `json:"physio_id"`
	AnatomyId uint32 `json:"anatomy_id"`

	//POPULATED
	Anatomy *Anatomy `gorm:"foreignKey:anatomy_id" json:"anatomiczone,omitempty"`
	Physio  *Physio  `gorm:"foreignKey:physio_id" json:"move,omitempty"`
}

func (c *LinkPhysioAnatomy) GetId() uint32 {
	return c.ID
}

func (c *LinkPhysioAnatomy) TableName() string {
	return "api_content.link_physio_anatomy"
}
