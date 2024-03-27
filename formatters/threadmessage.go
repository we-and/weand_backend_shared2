package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"time"
)

type ChatmessageFormattedItem struct {
	CreatedAt    time.Time `json:"created_at"`
	Content      string    `json:"content"`
	IsSentByUser bool      `json:"is_sent_by_user"`

	Pid           uint32                `json:"pid"`
	ThreadPid     uint32                `json:"thread_pid"`
	Type          string                `json:"type"`
	Viewers       []PersonFormattedItem `json:"viewers"`
	Creator       PersonFormattedItem   `json:"creator"`
	IsSoftDeleted bool                  `json:"is_soft_deleted"`
}

func FormatViewersFromLinkChatmessageViewer(items []m.LinkChatmessageViewer, me structs.Me) []PersonFormattedItem {
	res := []PersonFormattedItem{}
	for _, v := range items {
		if v.Person != nil && v.Person.HasUser() {
			//	res = append(res, FormatUserIsYou(*(v.Person.User), me))
		}
	}

	return res
}
func FormatChatmessages(items []m.Chatmessage, me structs.Me) []ChatmessageFormattedItem {
	res := []ChatmessageFormattedItem{}
	for _, v := range items {
		res = append(res, FormatChatmessage(v, me))
	}
	return res
}
func FormatChatmessage(v m.Chatmessage, me structs.Me) ChatmessageFormattedItem {
	res := ChatmessageFormattedItem{
		CreatedAt:     v.CreatedAt,
		Pid:           publicid.Obfuscate32bit(v.ID),
		IsSentByUser:  me.CheckPersonId(v.CreatorPersonId),
		Type:          v.Type,
		Content:       v.Content,
		IsSoftDeleted: v.IsSoftDeleted,
		ThreadPid:     publicid.Obfuscate32bit(v.ThreadId),
	}
	if v.Viewers != nil {
		res.Viewers = FormatViewersFromLinkChatmessageViewer(*(v.Viewers), me)
	}
	if v.Creator != nil {
		//	res.Creator = FormatPerson(*(v.Creator), me)
	}
	if !res.IsSentByUser {
		res.Creator = PersonFormattedItem{
			User: &UserFormattedItem{
				FirstName: "Team",
				LastName:  "Epic Flex Quest",
				Profile: ProfileFormattedItem{
					DisplaynameMode: "Team Epic Flex Quest",
				},
			},
		}
	}

	return res
}
