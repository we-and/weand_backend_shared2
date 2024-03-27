package models

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type WorkoutProfileRaw struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`

	Conditions pq.Int64Array `gorm:"type:integer[]" json:"conditions"`
	Gear       pq.Int64Array `gorm:"type:integer[]" json:"gear"`
	Sports     pq.Int64Array `gorm:"type:integer[]" json:"sports"`
	Goals      pq.Int64Array `gorm:"type:integer[]" json:"goals"`
	Experience pq.Int64Array `gorm:"type:integer[]" json:"experience"`

	Injuries pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null" json:"injuries"` //map[uint32]bool
	Focus    pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null" json:"focus"`    //map[uint32]string
}

func (c *WorkoutProfileRaw) GetId() uint32 {
	return c.ID
}

func (c *WorkoutProfileRaw) TableName() string {
	return "api_progress.workout_profile_raw"
}
