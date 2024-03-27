package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"time"
)

type FeedFormattedItem struct {
	Type      string     `json:"type"`
	Subtype   string     `json:"subtype"`
	CreatedAt *time.Time `json:"created_at"`
	//	IsYou bool ""
	Pid uint32 `json:"pid"`

	TeamPid           uint32 `json:"team_pid"`
	UserPid           uint32 `json:"user_pid"`
	RelatedPid        uint32 `json:"related_pid"`
	OccurenceTimetamp int64  `json:"occurence_timestamp"`
}

func FormatFeeds(items []m.Activity, me structs.Me) []FeedFormattedItem {
	res := []FeedFormattedItem{}
	for _, v := range items {
		res = append(res, FormatFeed(v, me))
	}
	return res
}

func FormatFeed(item m.Activity, me structs.Me) FeedFormattedItem {

	res := FeedFormattedItem{
		Type:              item.Type_,
		Subtype:           item.Subtype,
		CreatedAt:         item.CreatedAt,
		TeamPid:           publicid.Obfuscate32bit(item.TeamId),
		UserPid:           publicid.Obfuscate32bit(item.UserId),
		RelatedPid:        publicid.Obfuscate32bit(item.RelatedId),
		OccurenceTimetamp: item.OccurenceId,
		Pid:               publicid.Obfuscate32bit(item.ID),
	}
	return res
}
