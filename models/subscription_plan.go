package models

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name,omitempty"`

	//POPULATED
	Pricings []Pricing `gorm:"foreignKey:plan_id" json:"pricings,omitempty"`
}

func (c *Plan) GetId() uint32 {
	return c.ID
}

func (c *Plan) TableName() string {
	return "api_subscription.plan"
}
