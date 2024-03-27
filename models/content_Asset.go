package models

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx          int    `json:"idx"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	IsActive     bool   `json:"is_active"`
	IsGear       bool   `json:"is_gear"`
	Type         string `json:"type"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Lighting     string `json:"lighting"`

	LinksTag   []LinkAssetTag `gorm:"foreignKey:asset_id" json:"tags,omitempty"`
	LinksModel []Model        `gorm:"foreignKey:asset_id" json:"models,omitempty"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Asset) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "ASSET_NAME" {
				return k.Text
			}
		}
	}
	return c.Name

}
func (c *Asset) GetId() uint32 {
	return c.ID
}

func (c *Asset) TableName() string {
	return "api_content.asset"
}
