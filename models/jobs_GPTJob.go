package models

import (
	"time"

	"gorm.io/gorm"
)

type GPTJob struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	UserId           uint32 `json:"user_id"`
	UserKey          string `json:"user_key"`
	BatchId          string `json:"batch_id"`
	WorkoutprofileId uint32 `json:"workoutprofile_id"`
	ProfilerawId     uint32 `json:"profileraw_id"`
	Status           string `json:"status"`
	Type             string `json:"type"`
	ResultId         uint32 `json:"result_id"`

	Result     *GPTGenerateResult `gorm:"foreignKey:result_id" json:"result,omitempty"`
	AIGen      *AIGeneration      `gorm:"foreignKey:job_id" json:"aigen,omitempty"`
	Chats      []GPTChat          `gorm:"foreignKey:job_id" json:"gptchat,omitempty"`
	NextResult *GPTNextResult     `gorm:"foreignKey:result_id" json:"result,omitempty"`
}

func (c *GPTJob) GetId() uint32 {
	return c.ID
}

func (c *GPTJob) TableName() string {
	return "api_jobs.gpt_job"
}
