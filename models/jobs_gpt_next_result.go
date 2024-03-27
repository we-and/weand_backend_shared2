package models

import (
	_ "time"

	"gorm.io/gorm"
)

type GPTNextResult struct {
	gorm.Model
	ID               uint32 `json:"id" gorm:"primaryKey"`
	JobId            uint32 `json:"job_id"`
	ProgramId        uint32 `json:"program_id"`
	FirstworkoutId   uint32 `json:"firstworkout_id"`
	BatchId string `json:"batch_id"`
	Job     *GPTJob  `gorm:"foreignKey:job_id"`
	Program *Program `gorm:"foreignKey:program_id"`
}

func (c *GPTNextResult) GetId() uint32 {
	return c.ID
}

func (c *GPTNextResult) TableName() string {
	return "api_jobs.gpt_next_result"
}
