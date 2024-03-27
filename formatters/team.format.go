package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"time"
)

type TeamFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Pid       uint32     `json:"pid"`
	Level     int16      `json:"level"`

	Name        string `json:"name"`
	ColorAccent string `json:"color_accent"`
	ColorMain   string `json:"color_main"`
	CanEdit     bool   `json:"can_edit"`

	AvatarUrl string `json:"avatar_url,omitempty"`
	//CoachInvitekey  string              `json:"coach_invitekey,omitempty"`
	//	MemberInvitekey string              `json:"member_invitekey,omitempty"`
	Invitekey string `json:"invitekey,omitempty"`
	NbInvites uint32 `json:"nb_invites"`
	NbCoaches uint32 `json:"nb_coaches"`
	NbMembers uint32 `json:"nb_members"`
	//Coaches   []UserFormattedItem `json:"coaches"`
	Persons []PersonFormattedItem `json:"persons"`
	Org     *TeamFormattedItem    `json:"org,omitempty"`

	Invites []InviteFormattedItem `json:"invites"`
	//	Members    []UserFormattedItem     `json:"members"`
	Categories []CategoryFormattedItem `json:"categories"`

	Parent *TeamFormattedItem `json:"parent"`
}

func FormatTeams(items []m.Team, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {
		res = append(res, FormatTeam(v, me))
	}
	return res
}

func FormatTeamsFromSeat(items []m.Seat, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, FormatTeam(*v.Team, me))
		}
	}
	return res
}

func FormatTeamsFromLinkTeamOrg(items []m.LinkTeamOrg, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, FormatTeam(*v.Team, me))
		}
	}
	return res
}

func FormatTeamsFromLinkEventTeam(items []m.LinkEventTeam, me structs.Me) []TeamFormattedItem {
	res := []TeamFormattedItem{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, FormatTeam(*v.Team, me))
		}
	}
	return res
}

func FormatTeamPidsFromLinkEventTeam(items []m.LinkEventTeam) []uint32 {
	res := []uint32{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, publicid.Obfuscate32bit(v.TeamId))
		}
	}
	return res
}
func FormatTeamPidsFromLinkPollTeam(items []m.LinkPollTeam) []uint32 {
	res := []uint32{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, publicid.Obfuscate32bit(v.TeamId))
		}
	}
	return res
}
func FormatTeamPidsFromLinkAnnouncementTeam(items []m.LinkAnnouncementTeam) []uint32 {
	res := []uint32{}
	for _, v := range items {
		if v.Team != nil {
			res = append(res, publicid.Obfuscate32bit(v.TeamId))
		}
	}
	return res
}
func FormatTeam(item m.Team, me structs.Me) TeamFormattedItem {

	res := TeamFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		DeletedAt: item.DeletedAt,
		Name:      item.Name,
		CanEdit:   item.CanEdit(me),
		Level:     item.Level,

		NbInvites:   item.NbInvites,
		NbCoaches:   item.NbCoaches,
		AvatarUrl:   item.AvatarUrl,
		ColorAccent: item.ColorAccent,
		ColorMain:   item.ColorMain,
		Invitekey:   item.Invitekey,
		NbMembers:   item.NbMembers,

		Invites: FormatInvites(item.Invites, me),
		Persons: FormatPersonsFromSeats(item.Seats, me, item.ID),

		//		Coaches:    FormatUsersFromLinkByType(item.LinkUsers, "COACH", userId),
		//		Members:    FormatUsersFromLinkByType(item.LinkUsers, "MEMBER", userId),
		Categories: FormatCategoriesFromLinkTeam(item.LinkCategories),
	}
	if item.Parent != nil {
		parent := FormatTeam(*item.Parent, me)
		res.Parent = &parent
	}

	if item.LinkOrg != nil && item.LinkOrg.Org != nil {
		org := FormatOrgIsYou(*(*item.LinkOrg).Org, me)
		res.Org = &org
	}

	return res
}
