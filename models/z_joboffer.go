package models

import (
	"time"

	"gorm.io/gorm"
)

type Joboffer struct {
	gorm.Model
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	TitleFr   string     `json:"title_fr"`
	ContentFR string     `json:"content_fr"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Deleted   bool       `json:"deleted"`
	Active    bool       `json:"active"`
}

func (c *Joboffer) GetId() int {
	return c.ID
}

func (c *Joboffer) TableName() string {
	return "api_home.joboffer"
}
