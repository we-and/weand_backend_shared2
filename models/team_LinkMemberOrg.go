package models

import (
	"time"
)

type LinkMemberOrg struct {
	ID        uint32     `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserId uint32 `json:"user_id"`
	OrgId  uint32 `json:"org_id"`

	//POPULATED
	Member *User `gorm:"foreignKey:user_id",json:"member,omitempty"`
	Org    *Team `gorm:"foreignKey:org_id",json:"org,omitempty"`
}

func (c *LinkMemberOrg) GetId() uint32 {
	return c.ID
}

func (c *LinkMemberOrg) TableName() string {
	return "api_team.link_member_org"
}
