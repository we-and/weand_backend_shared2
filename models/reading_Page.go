package models

import (
	"time"

	"gorm.io/gorm"
)

type Page struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	ArticleId uint32 `json:"article_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Idx       uint32 `json:"idx"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Page) GetId() uint32 {
	return c.ID
}
func (c *Page) TableName() string {
	return "api_reading.page"
}

func (c *Page) GetContent(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "PAGE_CONTENT" {
				return k.Text
			}
		}
	}
	return c.Content
}
func (c *Page) GetTitle(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "PAGE_TITLE" {
				return k.Text
			}
		}
	}
	return c.Title
}
