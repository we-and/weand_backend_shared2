package models

import (
	_ "time"

	"gorm.io/gorm"
)

type GPTGenerateResult struct {
	gorm.Model
	ID               uint32 `json:"id" gorm:"primaryKey"`
	ProgramDesc      string `json:"program_desc"`
	Userprofile      string `json:"userprofile"`
	ProgramGuidance  string `json:"program_guidance"`
	ProgramName      string `json:"program_name"`
	ProgramMeaning   string `json:"program_meaning"`
	JobId            uint32 `json:"job_id"`
	ProgramId        uint32 `json:"program_id"`
	FirstworkoutId   uint32 `json:"firstworkout_id"`
	WorkoutprofileId uint32 `json:"workoutprofile_id"`

	Job     *GPTJob  `gorm:"foreignKey:job_id"`
	Program *Program `gorm:"foreignKey:program_id"`
}

func (c *GPTGenerateResult) GetId() uint32 {
	return c.ID
}

func (c *GPTGenerateResult) TableName() string {
	return "api_jobs.gpt_generate_result"
}
