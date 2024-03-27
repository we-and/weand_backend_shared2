package models

import (
	"time"

	"gorm.io/gorm"
)

type Goal struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx  uint32 `json:"idx"`
	Name string `json:"name"`
	Key  string `json:"key"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Goal) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "GOAL" {
				return k.Text
			}
		}
	}
	return c.Name

}

func (c *Goal) GetId() uint32 {
	return c.ID
}

func (c *Goal) TableName() string {
	return "api_content.goal"
}
