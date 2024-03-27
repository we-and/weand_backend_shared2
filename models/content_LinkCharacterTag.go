package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkAssetTag struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx            uint32 `json:"idx"` 
	TagId            uint32 `json:"tag_id"`
	AssetId uint32 `json:"asset_id"`
	
	//POPULATED
	Asset *Asset `gorm:"foreignKey:asset_id" json:"asset,omitempty"`
	Tag         *Tag         `gorm:"foreignKey:tag_id" json:"tsg,omitempty"`
}

func (c *LinkAssetTag) GetId() uint32 {
	return c.ID
}

func (c *LinkAssetTag) TableName() string {
	return "api_content.link_asset_tag"
}
