package formatters

import (
	models "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	"stretches-common-api/timezone"
	"time"
)

type AnnouncementReplyFormattedItem struct {
	CreatedAt       *time.Time `json:"created_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	Pid             uint32     `json:"pid"`
	PersonPid       uint32     `json:"person_pid"`
	AnnouncementPid uint32     `json:"poll_pid"`

	IsYou        bool                      `json:"is_you,omitempty"`
	Person       PersonFormattedItem       `json:"person"`
	Announcement AnnouncementFormattedItem `json:"poll"`
}

func FormatAnnouncementReplies(items []models.LinkAnnouncementPerson, me structs.Me, reqTzData *timezone.TzData) []AnnouncementReplyFormattedItem {
	res := []AnnouncementReplyFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAnnouncementReply(v, me, reqTzData))
	}
	return res
}

/*
	func FormatRepliesFromLinkAnnouncementReply(items []Announcementmodel.LinkAnnouncementUser, userId uint32) []ReplyFormattedItem {
		res := []ReplyFormattedItem{}
		for _, v := range items {
			if v.Team != nil {
				res = append(res, FormatReply(*v.Team, userId))
			}
		}
		return res
	}
*/
func FormatAnnouncementReply(item models.LinkAnnouncementPerson, me structs.Me, reqTzData *timezone.TzData) AnnouncementReplyFormattedItem {

	res := AnnouncementReplyFormattedItem{
		Pid:             publicid.Obfuscate32bit(item.ID),
		CreatedAt:       item.CreatedAt,
		UpdatedAt:       item.UpdatedAt,
		DeletedAt:       item.DeletedAt,
		PersonPid:       publicid.Obfuscate32bit(item.PersonId),
		AnnouncementPid: publicid.Obfuscate32bit(item.AnnouncementId),
		IsYou:           me.CheckPersonId(item.PersonId),
	}

	if item.Person != nil {
		user := FormatPersonUser(*(item.Person), me, false)
		res.Person = user
	}
	if item.Announcement != nil {
		Announcement := FormatAnnouncement(*(item.Announcement), me, reqTzData)
		res.Announcement = Announcement
	}

	return res
}
