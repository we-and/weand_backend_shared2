package models

import (
	"time"

	"gorm.io/gorm"
)

type Anatomy struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Idx      uint32  `json:"idx"`
	Name     string  `json:"name"`
	Level    uint32  `json:"level"`
	ParentId uint32  `json:"parent_id"`
	Posx     float32 `json:"posx"`
	Posy     float32 `json:"posy"`
	Labelx   float32 `json:"labelx"`
	Labely   float32 `json:"labely"`

	AnatomyFocus []Anatomyfocus `gorm:"foreignKey:anatomy_id" json:"focus,omitempty"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Anatomy) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "ZONE_NAME" {
				return k.Text
			}
		}
	}
	return c.Name

}
func (c *Anatomy) GetId() uint32 {
	return c.ID
}

func (c *Anatomy) TableName() string {
	return "api_content.anatomy"
}
