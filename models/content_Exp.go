package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx uint32 `json:"idx"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Key string `json:"key"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}


func (c *Experience) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "EXP_NAME" {
				return k.Text
			}
		}
	}
	return c.Name

}


func (c *Experience) GetDesc(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "EXP_DESC" {
				return k.Text
			}
		}
	}
	return c.Name

}

func (c *Experience) GetId() uint32 {
	return c.ID
}

func (c *Experience) TableName() string {
	return "api_content.experience"
}
