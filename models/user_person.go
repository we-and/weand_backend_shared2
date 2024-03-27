package models

import (
	structs "stretches-common-api/structs"
	"time"
	_ "time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UserId               *uint32              `json:"user_id"`
	UnregisteredId       *uint32              `json:"unregistered_id"`
	Notificationsettings NotificationSettings `json:"notifications_settings,omitempty"`
	Details              structs.UserDetails  `json:"details"`

	///POPULATED
	LinkTags     []LinkPersonTag `gorm:"foreignKey:person_id" json:"linktags,omitempty"`
	Unregistered *Unregistered   `gorm:"foreignKey:unregistered_id" json:"unregistered,omitempty"`
	User         *User           `gorm:"foreignKey:user_id" json:"user,omitempty"`
	UserKey      string          `json:"user_key,omitempty"`
	UserRights   Rights          `gorm:"foreignKey:user_id" json:"rights,omitempty"`
}

func (c *Person) GetRSVPPrefs() []string {
	notif := []string{}
	if c.Notificationsettings.Email.Rsvp {
		notif = append(notif, "EMAIL")
	}
	if c.Notificationsettings.Notif.Rsvp {
		notif = append(notif, "NOTIFICATION")
	}
	if c.Notificationsettings.SMS.Rsvp {
		notif = append(notif, "SMS")
	}
	return notif
}
func (c *Person) HasUser() bool {
	return c.UserId != nil && c.User != nil
}
func (c *Person) HasUnregistered() bool {
	return c.UnregisteredId != nil && c.Unregistered != nil
}

func (c *Person) IsMe(meUserId uint32) bool {
	return c.IsRegistered() && *(c.UserId) == meUserId
}
func (c *Person) GetUserId() uint32 {
	if c.UnregisteredId != nil {
		return *(c.UnregisteredId)
	}
	if c.UserId != nil {
		return *(c.UserId)
	}
	return 0
}
func (c *Person) IsRegistered() bool {
	if c.UnregisteredId != nil {
		return false
	}
	if c.UserId != nil {
		return true
	}
	return false
}
func (c *Person) GetId() uint32 {
	return c.ID
}

func (c *Person) TableName() string {
	return "api_user.person"
}
