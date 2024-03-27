package models

import (
	_ "time"

	"gorm.io/gorm"
)

type LinkThreadMember struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	ThreadId uint32 `json:"thread_id"`
	PersonId uint32 `json:"person_id"`

	//POPULATED FIELDS
	Person *Person `gorm:"foreignKey:person_id" json:"person,omitempty"`
}

func (c *LinkThreadMember) GetId() uint32 {
	return c.ID
}

func (c *LinkThreadMember) TableName() string {
	return "api_chat.link_thread_user_members"
}
