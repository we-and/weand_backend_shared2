package models

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	BatchId string `json:"batch_id"`

	ArenaKey     string     `json:"arena_key"`
	CoachKey     string     `json:"coach_key"`
	IntroMoveId  uint32     `json:"intro_move_id"`
	PlangroupId  uint32     `json:"plangroup_id"`
	ProgramId    uint32     `json:"program_id"`
	WeekNb       uint32     `json:"week_nb"`
	Dayoftheweek uint32     `json:"dayoftheweek"`
	Idx          uint32     `json:"idx"`
	WorkoutId    uint32     `json:"workout_id"`
	ScheduledFor *time.Time `json:"scheduled_for,omitempty"`

	//POPULATED
	
	IntroMove Move `gorm:"foreignKey:intro_move_id"`
	
	Plangroup *Plangroup `gorm:"foreignKey:plangroup_id" json:"plangroup,omitempty"`
	Tasks     []Task     `gorm:"foreignKey:lesson_id" json:"tasks,omitempty"`
	Workout   *Workout   `gorm:"foreignKey:workout_id" json:"workout,omitempty"`
}

func (c *Lesson) GetId() uint32 {
	return c.ID
}

func (c *Lesson) TableName() string {
	return "api_content.plan"
}
