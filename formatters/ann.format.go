package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	timezone "stretches-common-api/timezone"
	"time"
)

type AnnouncementFormattedItem struct {
	CreatedAt *time.Time           `json:"created_at"`
	DeletedAt *time.Time           `json:"deleted_at,omitempty"`
	UpdatedAt *time.Time           `json:"updated_at,omitempty"`
	Pid       uint32               `json:"pid"`
	Title     string               `json:"title"`
	Status    string               `json:"status"`
	Content   string               `json:"content"`
	CanEdit   bool                 `json:"can_edit"`
	Batches   []BatchFormattedItem `json:"batches"`
	Teams     []TeamFormattedItem  `json:"teams,omitempty"`
	TeamPids  []uint32             `json:"team_pids,omitempty"`
}

func FormatAnnouncementsFromLinkTeam(items []m.LinkAnnouncementTeam, me structs.Me, reqTzData *timezone.TzData) []AnnouncementFormattedItem {
	res := []AnnouncementFormattedItem{}
	for _, v := range items {
		if v.Announcement != nil {
			res = append(res, FormatAnnouncement(*(v.Announcement), me, reqTzData))
		}
	}
	return res
}
func FormatAnnouncements(items []m.Announcement, me structs.Me, reqTzData *timezone.TzData) []AnnouncementFormattedItem {
	res := []AnnouncementFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAnnouncement(v, me, reqTzData))
	}
	return res
}
func FormatAnnouncement(item m.Announcement, me structs.Me, reqTzData *timezone.TzData) AnnouncementFormattedItem {

	res := AnnouncementFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,

		DeletedAt: item.DeletedAt,
		CanEdit:   item.CanEdit(me),
		Status:    item.Status,
		Title:     item.Title,
		Content:   item.Content,
	}
	res.TeamPids = FormatTeamPidsFromLinkAnnouncementTeam(item.LinkTeams)
	res.Batches = FormatBatches(item.Batches, me, reqTzData)
	return res
}
