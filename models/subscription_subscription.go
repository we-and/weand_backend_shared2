package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	PricingId uint32     `json:"pricing_id,omitempty"`
	UserId    uint32     `json:"user_id,omitempty"`
	PlanId    uint32     `json:"plan_id,omitempty"`
	UserKey   string     `json:"user_key,omitempty"`
	Details   string     `json:"details,omitempty"`

	//POPULATED
	Pricing Pricing `gorm:"foreignKey:pricing_id" json:"pricing,omitempty"`
}

func (c *Subscription) GetId() uint32 {
	return c.ID
}

func (c *Subscription) TableName() string {
	return "api_subscription.subscription"
}
