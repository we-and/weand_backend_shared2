package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint32 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Type string `json:"type"`

	IconUrl  string `json:"icon_url"`
	Level    uint32 `json:"level"`
	ParentId uint32 `json:"parent_id"`
	//POPULATED
	Parent *Category  `gorm:"foreignKey:parent_id" json:"parent,omitempty"`
	Sub    []Category `gorm:"foreignKey:parent_id" json:"sub,omitempty"`
}

func (c *Category) GetId() uint32 {
	return c.ID
}

func (c *Category) TableName() string {
	return "api_internal.category"
}
