package models

import (
	"gorm.io/gorm"

	"time"
	_ "time"
)

type HealthProfile struct {
	gorm.Model
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//Deleted   bool       `json:"deleted,omitempty"`

	UserId uint32 `json:"user_id,omitempty"`
	Name   string `json:"name"`

	//POPULATED FIELDS
	User User `gorm:"foreignKey:user_id",json:"user,omitempty"`
}

func (c *HealthProfile) GetId() uint32 {
	return c.ID
}

func (c *HealthProfile) TableName() string {
	return "api_profile.healthprofile"
}
