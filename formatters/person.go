package formatters

import (
	m "stretches-common-api/models"
	structs "stretches-common-api/structs"

	"stretches-common-api/publicid"
)

type PersonFormattedItem struct {
	Identifier      string              `json:"identifier,omitempty"`
	TeamRole        string              `json:"team_role,omitempty"`
	IsYou           bool                `json:"is_you,omitempty"`
	IsRegistered    bool                `json:"is_registered,omitempty"`
	Pid             uint32              `json:"pid"`
	TeamPid         *uint32             `json:"team_pid"`
	RSVPPreferences []string            `json:"rsvp_pref"`
	Tags            []PersontagFormattedItem  `json:"tags,omitempty"`
	Details         structs.UserDetails `json:"details"`

	User         *UserFormattedItem         `json:"user,omitempty"`
	Team         *TeamFormattedItem         `json:"team,omitempty"`
	Unregistered *UnregisteredFormattedItem `json:"unregistered,omitempty"`
}

func FormatPerson(p m.Person, me structs.Me, teamId uint32, isForEditor bool) PersonFormattedItem {
	var res PersonFormattedItem
	if p.HasUser() {
		res = FormatPersonUser(p, me, isForEditor)
		res.IsRegistered = true
	} else if p.HasUnregistered() {
		res = FormatPersonUnregistered(p, me, isForEditor)
		res.IsRegistered = false
	}
	if isForEditor {
		res.RSVPPreferences = p.GetRSVPPrefs()
		res.Tags = FormatPersontagsFromLink(p.LinkTags, teamId)
	}
	res.Details = p.Details
	return res
}
func FormatPersonsAndTeamFromSeats(seats []m.Seat, me structs.Me, teamId uint32) []PersonFormattedItem {
	res := []PersonFormattedItem{}
	isEditor := false
	for _, v := range seats {
		//if user
		if v.HasUser() && me.CheckPersonId((v.PersonId)) {
			isEditor = v.IsEditor()
		}
	}

	for _, s := range seats {
		if s.HasPerson() {
			p := *(s.Person)
			formatted := FormatPerson(p, me, teamId, isEditor)
			formatted.TeamRole = s.Type
		//	t := FormatTeam(*(s.Team), me)
		//	formatted.Team = &t
			if len(formatted.TeamRole) == 0 {
				formatted.TeamRole = "MEMBER"
			}

			res = append(res, formatted)
		}
	}
	return res
}
func FormatPersonsFromSeats(seats []m.Seat, me structs.Me, teamId uint32) []PersonFormattedItem {
	res := []PersonFormattedItem{}
	isEditor := false
	for _, v := range seats {
		//if user
		if v.HasUser() && me.CheckPersonId((v.PersonId)) {
			isEditor = v.IsEditor()
		}
	}

	for _, s := range seats {
		if s.HasPerson() {
			p := *(s.Person)
			formatted := FormatPerson(p, me, teamId, isEditor)
			formatted.TeamRole = s.Type
			if len(formatted.TeamRole) == 0 {
				formatted.TeamRole = "MEMBER"
			}

			res = append(res, formatted)
		}
	}
	return res
}

func FormatPersonUser(p m.Person, me structs.Me, isForEditor bool) PersonFormattedItem {
	u := p.User

	res := PersonFormattedItem{
		IsRegistered:    true,
		RSVPPreferences: p.GetRSVPPrefs(),

		Pid:   publicid.Obfuscate32bit(p.ID),
		IsYou: me.CheckUserId(u.ID),
	}
	if p.User != nil {
		formatted := FormatUser(*(p.User), me)
		formatted.PersonPid = publicid.Obfuscate32bit(p.ID)
		res.User = &formatted
	}
	return res
}

func FormatPersonsFromLinkCoachOrg(links []m.Seat, me structs.Me, teamId uint32, isForEditor bool) []PersonFormattedItem {
	res := []PersonFormattedItem{}
	for _, v := range links {
		if v.Person != nil && v.Type == "COACH" {
			res = append(res, FormatPerson(*v.Person, me, teamId, isForEditor))
		}
	}
	return res
}
func FormatPersonsFromLinkByType(links []m.Seat, memberType string, me structs.Me, teamId uint32, isForEditor bool) []PersonFormattedItem {
	res := []PersonFormattedItem{}
	for _, v := range links {
		if v.Person != nil {
			if v.Type == memberType {
				res = append(res, FormatPerson(*v.Person, me, teamId, isForEditor))
			}
		}
	}
	return res
}
func FormatPersonUnregistered(p m.Person, me structs.Me, isForEditor bool) PersonFormattedItem {

	res := PersonFormattedItem{
		IsRegistered: false,
		Pid:          publicid.Obfuscate32bit(p.ID),
		IsYou:        me.CheckPersonId(p.ID),
	}
	if p.Unregistered != nil {
		u := FormatUnregistered(*(p.Unregistered), me, isForEditor)
		u.PersonPid = publicid.Obfuscate32bit(p.ID)
		res.Unregistered = &u

	}
	return res
}
