package models

import "gorm.io/gorm"

type Persontag struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	ParentId uint32 `json:"parent_id"`
	Level    uint32 `json:"level"`
	Name     string `json:"name"`

	IconUrl string `json:"icon_url"`
	//POPULATED
	//POPULATED
	Parent *Persontag  `gorm:"foreignKey:parent_id" json:"parent,omitempty"`
	Sub    []Persontag `gorm:"foreignKey:parent_id" json:"sub,omitempty"`
}

func (c *Persontag) GetId() uint32 {
	return c.ID
}

func (c *Persontag) TableName() string {
	return "api_internal.user_tag"
}
