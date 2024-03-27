package models

import (
	"time"
)

type Seat struct {
	ID        uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Type      string     `json:"type"`
	PersonId  uint32     `json:"person_id"`
	UserId    *uint32    `json:"user_id"`
	TeamId    uint32     `json:"team_id"`

	//POPULATED
	Person *Person `gorm:"foreignKey:person_id" json:"person,omitempty"`
	Team   *Team   `gorm:"foreignKey:team_id" json:"team,omitempty"`
}

func (c *Seat) GetId() uint32 {
	return c.ID
}
func (c *Seat) HasPerson() bool {
	return c.PersonId > 0 && c.Person != nil
}
func (c *Seat) HasUser() bool {
	return c.UserId != nil && c.Person != nil && (*(c.Person)).User != nil
}

func (c *Seat) TableName() string {
	return "api_team.seat"
}

func (c *Seat) IsEditor() bool {
	return c.Type == "COACH" || c.Type == "SUPPORT"
}
