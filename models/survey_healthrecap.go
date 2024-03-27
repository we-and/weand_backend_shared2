package models

import (
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Healthrecap struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`
	//	CreatedAt *time.Time `json:"created_at,omitempty"`
	//	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Goals      pq.Int64Array `gorm:"type:integer[]" json:"goals,omitempty"`
	UserKey    string        `json:"user_key,omitempty"`
	UserId     uint32        `json:"user_id,omitempty"`
	Experience pq.Int64Array `gorm:"type:integer[]" json:"experience,omitempty"`
	Conditions pq.Int64Array `gorm:"type:integer[]"  json:"conditions,omitempty"`
	//
	InjuryManagement pgtype.JSONB `gorm:"type:jsonb;default:'{}';not null"`
	//
	Injuries pq.Int64Array `gorm:"type:integer[]"`
}

func (c *Healthrecap) GetId() uint32 {
	return c.ID
}

func (c *Healthrecap) TableName() string {
	return "api_survey.health_recap"
}
