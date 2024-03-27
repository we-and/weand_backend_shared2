package models

import (
	"time"

	"gorm.io/gorm"
)

type Pricing struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	PlanId    uint32     `json:"plan_id,omitempty"`
	Price     uint32     `json:"price"`
	Frequency string     `json:"frequency"`
	Currency  string     `json:"currency"`

	//POPULATED
	Plan Plan `gorm:"foreignKey:plan_id" json:"plan,omitempty"`
}

func (c *Pricing) GetId() uint32 {
	return c.ID
}

func (c *Pricing) TableName() string {
	return "api_subscription.pricing"
}
