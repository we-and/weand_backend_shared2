package models

import (
	"stretches-common-api/structs"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	ID        uint32     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Title   string `json:"title"`
	Type    string `json:"type"`
	Status  string `json:"status"`
	Choices string `json:"choices"`
	//POPULATED FIELDS
	LinkTeams []LinkPollTeam   `gorm:"foreignKey:poll_id" json:"linkteams,omitempty"`
	Replies   []LinkPollPerson `gorm:"foreignKey:poll_id" json:"replies,omitempty"`
	Batches []SendBatch `gorm:"foreignKey:related_id" json:"batches,omitempty"`

}

func (c *Poll) GetId() uint32 {
	return c.ID
}
func (c *Poll) GetChoicesArray() []string {
	if len(c.Choices) == 0 {
		return []string{}
	}
	return strings.Split(c.Choices, "###")

}
func (c *Poll) SetChoicesArray(choices []string) {
	c.Choices = strings.Join(choices, "###")
}

func (c *Poll) TableName() string {
	return "api_poll.poll"
}

func (c *Poll) CanEdit(me structs.Me) bool {
	for _, liTeam := range c.LinkTeams {
		if liTeam.Team != nil {
			if liTeam.Team.CanEdit(me) {
				return true
			}
		}
	}
	return false
}
