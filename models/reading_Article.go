package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Title       string `json:"title"`
	Content     string `json:"content"`
	DurationMin int    `json:"duration_min"`
	NbViews     int    `json:"nb_views"`
	Pages       []Page `gorm:"foreignKey:article_id;references:id" json:"pages,omitempty"`

	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Article) GetId() uint32 {
	return c.ID
}
func (c *Article) TableName() string {
	return "api_reading.article"
}

func (c *Article) GetContent(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "ARTICLE_CONTENT" {
				return k.Text
			}
		}
	}
	return c.Content
}
func (c *Article) GetTitle(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "ARTICLE_TITLE" {
				return k.Text
			}
		}
	}
	return c.Title
}
