package models

import (
	"time"

	"gorm.io/gorm"
)

type AIGeneration struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	CompletionTokens int    `json:"completion_tokens"`
	ElapsedTimeSec   int    `json:"elapsed_time_sec"`
	PromptTokens     int    `json:"prompt_tokens"`
	JobId            uint32 `json:"job_id"`
	Type             string `json:"type"`
	BatchId          string `json:"batch_id"`
	UserProfile      string `json:"user_profile"`
	Req              string `json:"req"`
}

func (c *AIGeneration) GetId() uint32 {
	return c.ID
}

func (c *AIGeneration) TableName() string {
	return "api_progress.ai_generation"
}
