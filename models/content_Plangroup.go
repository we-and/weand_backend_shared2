package models

import (
	"time"

	"gorm.io/gorm"
)

type Plangroup struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	BatchId string `json:"batch_id"`

	ProgramId uint32 `json:"program_id"`
	Idx       uint32 `json:"idx"`
	Name      string `json:"name"`

	//POPULATED
	Plans []Lesson `gorm:"foreignKey:plangroup_id" json:"plans,omitempty"`
}

func (c *Plangroup) GetId() uint32 {
	return c.ID
}

func (c *Plangroup) TableName() string {
	return "api_content.plangroup"
}
