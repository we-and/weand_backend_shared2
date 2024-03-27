package models

import (
	"gorm.io/gorm"

	"time"
	_ "time"
)

type Profile struct {
	gorm.Model
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//Deleted   bool       `json:"deleted,omitempty"`

	UserId uint32 `json:"user_id,omitempty"`
	Name   string `json:"name"`
	Key    string `json:"key"`
	//	AvatarIcon string `json:"avatar_icon"`
	AvatarUrl       string `json:"avatar_url"`
	DisplaynameMode string `json:"displayname_mode"`
	CoverUrl        string `json:"cover_url"`

	//POPULATED FIELDS
	User User `gorm:"foreignKey:user_id" json:"user,omitempty"`
}

func (c *Profile) GetId() uint32 {
	return c.ID
}

func (c *Profile) TableName() string {
	return "api_user.profile"
}
