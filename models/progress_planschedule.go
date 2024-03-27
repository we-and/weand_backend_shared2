package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanSchedule struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`


	ProgramId    uint32 `json:"program_id"`
	ScheduledForStr string `json:"scheduled_for_str,omitempty"`

	//POPULATED
}

func (c *PlanSchedule) GetId() uint32 {
	return c.ID
}

func (c *PlanSchedule) TableName() string {
	return "api_progress.plan_schedule"
}
