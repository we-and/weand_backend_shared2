package models

import (
	"time"

	"gorm.io/gorm"
)

type Sport struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx uint32 `json:"idx"`
	ParentId uint32 `json:"parent_id"`
	Name string `json:"name"`
	Key string `json:"key"`
	Level int `json:"level"`
	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Sport) GetId() uint32 {
	return c.ID
}


func (c *Sport) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "SPORT" {
				return k.Text
			}
		}
	}
	return c.Name

}

func (c *Sport) TableName() string {
	return "api_content.sport"
}
