package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Name     string `json:"name"`
	Type     string `json:"type"`
	Key     string `json:"desc"`
	
	LinksAsset []LinkAssetTag `gorm:"foreignKey:tag_id" json:"asset,omitempty"`
}

func (c *Tag) GetId() uint32 {
	return c.ID
}

func (c *Tag) TableName() string {
	return "api_content.tag"
}
