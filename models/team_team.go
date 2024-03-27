package models

import (
	"fmt"
	"time"

	"stretches-common-api/structs"
	util "stretches-common-api/utils"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	ID              uint32     `json:"id" gorm:"primaryKey"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	CoachInvitekey  string     `json:"coach_invitekey"`
	MemberInvitekey string     `json:"member_invitekey"`
	Invitekey       string     `json:"invitekey"`
	Name            string     `json:"name"`
	ColorAccent     string     `json:"color_accent"`
	ColorMain       string     `json:"color_main"`
	AvatarUrl       string     `json:"avatar_url"`
	Level           int16      `json:"level"`
	ParentId        uint32     `json:"parent_id"`
	NbInvites       uint32     `json:"nb_invites"`
	NbCoaches       uint32     `json:"nb_coaches"`
	NbMembers       uint32     `json:"nb_members"`

	//POPULATED
	LinkSubteams []LinkTeamOrg `gorm:"foreignKey:org_id" json:"linkteams,omitempty"`
	//	LinkUsers []Seat `gorm:"foreignKey:org_id" json:"linkusers,omitempty"`

	LinkEvents     []LinkEventTeam    `gorm:"foreignKey:team_id" json:"linkevents,omitempty"`
	Seats          []Seat             `gorm:"foreignKey:team_id" json:"seats,omitempty"`
	LinkCategories []LinkTeamCategory `gorm:"foreignKey:team_id" json:"linkmembers,omitempty"`
	LinkOrg        *LinkTeamOrg       `gorm:"foreignKey:team_id" json:"linkorg,omitempty"`
	Invites        []Invite           `gorm:"foreignKey:team_id" json:"invites,omitempty"`
	Parent         *Team              `gorm:"foreignKey:parent_id" json:"parent,omitempty"`
	SendJobs       []SendJob          `gorm:"foreignKey:team_id" json:"jobs,omitempty"`
	SendBatch      []SendBatch        `gorm:"foreignKey:team_id" json:"batches,omitempty"`
}

func (c *Team) GetId() uint32 {
	return c.ID
}

func (c *Team) TableName() string {
	return "api_team.team"
}

func (c *Team) CanEdit(me structs.Me) bool {
	if c.Seats == nil {
		return false
	}
	for _, li := range c.Seats {
		if li.IsEditor() && (li.PersonId > 0) && ((li.PersonId) == me.PersonId) {
			return true
		}
	}
	return false
}
func (c *Team) GenerateInviteKey() string {
	key := fmt.Sprintf("%v%v%v%v-%v", util.RandomUpperConsonantString(1), util.RandomUpperVoyellString(1), util.RandomUpperConsonantString(1), util.RandomUpperVoyellString(1), util.RandomNumberString(2))
	//	key := fmt.Sprintf("%v-%v", util.RandomUpperAlphaString(3), util.RandomNumberString(3))
	c.Invitekey = key
	return key
}
