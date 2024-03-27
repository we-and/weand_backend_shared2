package formatters

import (
	models "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"stretches-common-api/timezone"
	"time"
)

type PollReplyFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Pid       uint32     `json:"pid"`
	PersonPid uint32     `json:"person_pid"`
	PollPid   uint32     `json:"poll_pid"`
	Reply     string     `json:"reply"`

	IsYou  bool                `json:"is_you,omitempty"`
	Person PersonFormattedItem `json:"person"`
	Poll   PollFormattedItem   `json:"poll"`
}

func FormatPollReplies(items []models.LinkPollPerson, me structs.Me, reqTzData *timezone.TzData) []PollReplyFormattedItem {
	res := []PollReplyFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPollReply(v, me, reqTzData))
	}
	return res
}

/*
	func FormatRepliesFromLinkPollReply(items []Pollmodel.LinkPollUser, userId uint32) []ReplyFormattedItem {
		res := []ReplyFormattedItem{}
		for _, v := range items {
			if v.Team != nil {
				res = append(res, FormatReply(*v.Team, userId))
			}
		}
		return res
	}
*/
func FormatPollReply(item models.LinkPollPerson, me structs.Me, reqTzData *timezone.TzData) PollReplyFormattedItem {

	res := PollReplyFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		DeletedAt: item.DeletedAt,
		Reply:     item.Reply,
		PersonPid: publicid.Obfuscate32bit(item.PersonId),
		PollPid:   publicid.Obfuscate32bit(item.PollId),
		IsYou:     me.CheckPersonId(item.PersonId),
	}

	if item.Person != nil {
		user := FormatPersonUser(*(item.Person), me, false)
		res.Person = user
	}
	if item.Poll != nil {
		Poll := FormatPoll(*(item.Poll), me, reqTzData)
		res.Poll = Poll
	}

	return res
}
