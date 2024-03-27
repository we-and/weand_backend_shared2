package models

import (
	"time"

	"gorm.io/gorm"
)

type GPTChat struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Role    string `json:"role"`
	Raw     string `json:"raw"`
	UserId  uint32 `json:"user_id"`
	UserKey string `json:"user_key"`
	JobId   uint32 `json:"job_id"`
	Content string `json:"content"`
}

func (c *GPTChat) GetId() uint32 {
	return c.ID
}

func (c *GPTChat) TableName() string {
	return "api_progress.gpt_chat"
}
