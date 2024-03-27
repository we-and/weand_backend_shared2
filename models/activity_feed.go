package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Type_       string `json:"type_,omitempty"`
	Subtype     string `json:"subtype,omitempty"`
	UserId      uint32 `json:"user_id,omitempty"`
	TeamId      uint32 `json:"team_id,omitempty"`
	RelatedId   uint32 `json:"related_id,omitempty"`
	OccurenceId int64  `json:"occurence_id,omitempty"`
}

func (c *Activity) GetId() uint32 {
	return c.ID
}

func (c *Activity) TableName() string {
	return "api_feed.activity"
}
