package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"time"
)

type ThreadFormattedItem struct {
	CreatedAt          time.Time                  `json:"created_at"`
	LastMessageId      int                        `json:"last_message_id"`
	Type               string                     `json:"type"`
	Name               string                     `json:"name"`
	Hashtag            string                     `json:"hashtag"`
	IsArchived         bool                       `json:"is_archived"`
	LastMessage        ChatmessageFormattedItem   `json:"last_message,omitempty"`
	LastMessageExtract string                     `json:"last_message_extract,omitempty"`
	Messages           []ChatmessageFormattedItem `json:"messages"`
	Pid                uint32                     `json:"pid"`
	//Id                 uint32                     `json:"id"`

	MembersCustom     []PersonFormattedItem `json:"members_custom"`
	MembersCustomPids []uint32              `json:"members_custom_pids"`
}

func FormatThreads(items []m.Thread, me structs.Me) []ThreadFormattedItem {
	res := []ThreadFormattedItem{}
	for _, v := range items {
		res = append(res, FormatThread(v, me))
	}
	return res
}
func FormatThreadsForAdmin(items []m.Thread, me structs.Me) []ThreadFormattedItem {
	res := []ThreadFormattedItem{}
	for _, v := range items {
		res = append(res, FormatThreadForAdmin(v, me))
	}
	return res
}
func FormatThread(v m.Thread, me structs.Me) ThreadFormattedItem {
	res := ThreadFormattedItem{
		CreatedAt:          v.CreatedAt,
		Hashtag:            v.Hashtag,
		Name:               v.Name,
		IsArchived:         v.IsArchived,
		Type:               v.Type,
		Messages:           FormatChatmessages(v.Messages, me),
		LastMessageExtract: v.LastMessageExtract,

		Pid: publicid.Obfuscate32bit(v.ID),
	}
	if v.LastMessage != nil {
		res.LastMessage = FormatChatmessage(*(v.LastMessage), me)
	}
	//	if v.CustomMembers != nil {
	//res.MembersCustomPids = FormatUsersPidsFromLinkThreadMember(v.CustomMembers)
	//res.MembersCustom = FormatUsersFromLinkThreadMember((v.CustomMembers), userId)
	//	}
	return res
}

func FormatThreadForAdmin(v m.Thread, me structs.Me) ThreadFormattedItem {
	res := ThreadFormattedItem{
		CreatedAt:          v.CreatedAt,
		Hashtag:            v.Hashtag,
		Name:               v.Name,
		//Id:                 v.ID,
		IsArchived:         v.IsArchived,
		Type:               v.Type,
		Messages:           FormatChatmessages(v.Messages, me),
		LastMessageExtract: v.LastMessageExtract,

		Pid: publicid.Obfuscate32bit(v.ID),
	}
	if v.LastMessage != nil {
		res.LastMessage = FormatChatmessage(*(v.LastMessage), me)
	}
	//	if v.CustomMembers != nil {
	//res.MembersCustomPids = FormatUsersPidsFromLinkThreadMember(v.CustomMembers)
	//res.MembersCustom = FormatUsersFromLinkThreadMember((v.CustomMembers), userId)
	//	}
	return res
}
