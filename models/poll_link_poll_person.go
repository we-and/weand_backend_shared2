package models

import (
	"stretches-common-api/utils"
	"strings"
	"time"
)

type LinkPollPerson struct {
	ID       uint32 `json:"id"`
	PollId   uint32 `json:"poll_id"`
	PersonId uint32 `json:"person_id"`
	Reply    string `json:"reply"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Person *Person `gorm:"foreignKey:person_id" json:"person,omitempty"`
	Poll   *Poll   `gorm:"foreignKey:poll_id" json:"poll,omitempty"`
}

func (c *LinkPollPerson) GetId() uint32 {
	return c.ID
}

func (c *LinkPollPerson) GetChoicesArray() []string {
	return strings.Split(c.Reply, "###")
}
func (c *LinkPollPerson) AddToReplyIfNotExisting(choice string) {
	rep := c.GetChoicesArray()
	alreadyExists := false
	for _, r := range rep {
		if r == choice {
			alreadyExists = true
			break
		}
	}
	if !alreadyExists {
		rep = append(rep, choice)
		c.SetChoicesArray(rep)
	}
}
func (c *LinkPollPerson) RemoveFromReplyIfExisting(choice string) {
	rep := c.GetChoicesArray()
	alreadyExists := -1
	for k, r := range rep {
		if r == choice {
			alreadyExists = k
			break
		}
	}
	if alreadyExists > -1 {
		rep = utils.RemoveFromStringArray(rep, alreadyExists)
		c.SetChoicesArray(rep)
	}
}
func (c *LinkPollPerson) SetChoicesArray(choices []string) {
	c.Reply = strings.Join(choices, "###")
}

func (c *LinkPollPerson) TableName() string {
	return "api_poll.link_poll_user"
}
