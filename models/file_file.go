package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID              uint32 `json:"id" gorm:"primaryKey"`
	FileUrl         string `json:"file_url"`
	Name            string `json:"name"`
	Status          string `json:"status"`
	Size            int64  `json:"size"`
	TeamId          uint32 `json:"team_id"`
	CreatorPersonId uint32 `json:"creator_person_id"`
	Extension       string `json:"extension"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Batches []SendBatch `gorm:"foreignKey:related_id" json:"batches,omitempty"`
}

func (c *File) GetId() uint32 {
	return c.ID
}

func (c *File) TableName() string {
	return "api_file.file"
}
