package formatters

import (
	m "stretches-common-api/models"
	structs "stretches-common-api/structs"
	"time"
)

type OrgFormattedItem struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Pid       uint32     `json:"pid"`

	NbInvites uint32 `json:"nb_invites"`
	NbCoaches uint32 `json:"nb_coaches"`
	NbMembers uint32 `json:"nb_members"`
	Name      string `json:"name,omitempty"`
	Invitekey string `json:"invitekey,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	CoverUrl  string `json:"cover_url,omitempty"`

	Coaches []PersonFormattedItem `json:"coaches"`
	Members []PersonFormattedItem `json:"members"`
	Invites []InviteFormattedItem `json:"invites"`

	Teams []TeamFormattedItem `json:"teams"`
}

func FormatOrgsIsYou(items []m.Team, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {

		res = append(res, FormatOrgIsYou(v, me))
	}
	return res
}
func FormatOrgasFromLinkCoachOrg(items []m.LinkUserOrg, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {
		if v.Org != nil {
			res = append(res, FormatOrgIsYou(*v.Org, me))
		}
	}
	return res
}

func FormatOrgIsYou(item m.Team, me structs.Me) TeamFormattedItem {
	return FormatTeam(item, me)
}
