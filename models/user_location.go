package models

import (
	"time"
	_ "time"
)

type Location struct {
	//	gorm.Model
	ID        uint32     `json:"id",gorm:"primaryKey"`
	Long      float64    `json:"long"`
	UserId    uint32     `json:"user_id"`
	Lat       float64    `json:"lat"`
	Radius    float64    `json:"radius"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *Location) GetId() uint32 {
	return c.ID
}

func (c *Location) TableName() string {
	return "api_user.location"
}
