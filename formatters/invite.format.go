package formatters

import (
	m "stretches-common-api/models"
	structs "stretches-common-api/structs"
	"time"
)

type InviteFormattedItem struct {
	CreatedAt *time.Time          `json:"created_at"`
	Pid       uint32              `json:"pid"`
	Name      string              `json:"name,omitempty"`
	Email     string              `json:"email,omitempty"`
	Phone     string              `json:"phone,omitempty"`
	InvitedBy PersonFormattedItem `json:"invitedby"`
}

func FormatInvites(items []m.Invite, me structs.Me) []InviteFormattedItem {
	res := []InviteFormattedItem{}
	for _, v := range items {
		res = append(res, FormatInvite(v, me))
	}
	return res
}
func FormatInvite(item m.Invite, me structs.Me) InviteFormattedItem {
	return InviteFormattedItem{

		CreatedAt: item.CreatedAt,
		Name:      item.Name,
		Email:     item.Email,
		Phone:     item.Phone,
		InvitedBy: FormatPerson(item.InvitedByPerson, me, 0, false),
	}
}
