package models

import (
	"time"
	_ "time"

	"gorm.io/gorm"
)

type Newsletter struct {
	gorm.Model
	ID        int        `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Deleted   bool       `json:"deleted,omitempty"`
}

func (c *Newsletter) GetId() int {
	return c.ID
}

func (c *Newsletter) TableName() string {
	return "api_user.newsletter"
}

func (c *Newsletter) SetDeleted() {
	c.Deleted = true
}
