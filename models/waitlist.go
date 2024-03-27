package models

import (
	"time"
	_ "time"
)

type Waitlist struct {
	//	gorm.Model
	ID          uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Deleted     bool       `json:"deleted,omitempty"`
	IsConfirmed bool       `json:"is_confirmed,omitempty"`

	//POPULATED FIELDS
	Email string `json:"email,omitempty"`
}

func (c *Waitlist) GetId() uint32 {
	return c.ID
}

func (c *Waitlist) TableName() string {
	return "api_user.waitlist"
}

func (c *Waitlist) SetDeleted() {
	c.Deleted = true
}
