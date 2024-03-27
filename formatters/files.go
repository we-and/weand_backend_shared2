package formatters

import (
	m "stretches-common-api/models"
	"time"

	"stretches-common-api/publicid"
)

type FileFormattedItem struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Pid       uint32     `json:"pid"`

	FileUrl   string `json:"file_url"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	Name      string `json:"name"`
	TeamType  string `json:"team_type"`
	TeamPid   uint32 `json:"team_pid"`
}

func FormatFilesFromLinks(items []m.LinkFileTeam) []FileFormattedItem {
	res := []FileFormattedItem{}
	for _, v := range items {
		if v.File != nil {
			loc := FormatFile(*(v.File))
			res = append(res, loc)

			//			loc.ParentPid= public.Obfuscate32bit(v.Teamlo),

		}
	}
	return res
}

func FormatFiles(items []m.File) []FileFormattedItem {
	res := []FileFormattedItem{}
	for _, v := range items {
		res = append(res, FormatFile(v))
	}
	return res
}
func FormatFile(item m.File) FileFormattedItem {

	return FileFormattedItem{
		Extension: item.Extension,
		Name:      item.Name,
		Size:      item.Size,
		FileUrl:   item.FileUrl,
		CreatedAt: item.CreatedAt,
		TeamPid:   publicid.Obfuscate32bit(item.TeamId),
		Pid:       publicid.Obfuscate32bit(item.ID),
	}
}
