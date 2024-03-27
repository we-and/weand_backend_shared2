package models

import (
	"time"
)

type Invite struct {
	//	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	InvitedBy uint32 `json:"invited_by"`
	OrgId     uint32 `json:"org_id"`
	TeamId    uint32 `json:"team_id"`
	Fulfilled bool   `json:"fulfilled"`

	InvitedByPerson Person `gorm:"foreignKey:invited_by" json:"invitedbyperson,omitempty"`
}

func (c *Invite) GetId() uint32 {
	return c.ID
}

func (c *Invite) TableName() string {
	return "api_user.invite"
}
