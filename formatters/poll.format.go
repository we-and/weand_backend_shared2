package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	timezone "stretches-common-api/timezone"
	"time"
)

type PollFormattedItem struct {
	CreatedAt *time.Time           `json:"created_at"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty"`
	UpdatedAt *time.Time           `json:"updated_at,omitempty"`
	Batches   []BatchFormattedItem `json:"batches"`
	Pid       uint32               `json:"pid"`
	Title     string               `json:"title"`
	Status    string               `json:"status"`
	Choices   []string             `json:"choices"`
	CanEdit   bool                 `json:"can_edit"`

	Teams    []TeamFormattedItem `json:"teams,omitempty"`
	TeamPids []uint32            `json:"team_pids,omitempty"`

	Replies []PollReplyFormattedItem `json:"replies"`
}

func FormatPollsFromLinkTeam(items []m.LinkPollTeam, me structs.Me, reqTzData *timezone.TzData) []PollFormattedItem {
	res := []PollFormattedItem{}
	for _, v := range items {
		if v.Poll != nil {
			res = append(res, FormatPoll(*(v.Poll), me, reqTzData))
		}
	}
	return res
}
func FormatPolls(items []m.Poll, me structs.Me, reqTzData *timezone.TzData) []PollFormattedItem {
	res := []PollFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPoll(v, me, reqTzData))
	}
	return res
}
func FormatPoll(item m.Poll, me structs.Me, reqTzData *timezone.TzData) PollFormattedItem {

	res := PollFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		DeletedAt: item.DeletedAt,
		CanEdit:   item.CanEdit(me),
		Status:    item.Status,
		Title:     item.Title,
		Choices:   item.GetChoicesArray(),
		Replies:   FormatPollReplies(item.Replies, me, reqTzData),
	}
	res.TeamPids = FormatTeamPidsFromLinkPollTeam(item.LinkTeams)
	res.Batches = FormatBatches(item.Batches, me, reqTzData)

	return res
}
