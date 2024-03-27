package models

import (
	"time"

	"gorm.io/gorm"
)

type Movegroup struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	BatchId   string         `json:"batch_id"`

	Name string `json:"name"`

	LinksMovegroupMove []LinkMovegroupMove `gorm:"foreignKey:movegroup_id" json:"links,omitempty"`
}

func (c *Movegroup) GetId() uint32 {
	return c.ID
}

func (c *Movegroup) TableName() string {
	return "api_content.movegroup"
}
