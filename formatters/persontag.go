package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type PersontagFormattedItem struct {
	//	CreatedAt *time.Time `json:"created_at,omitempty"`
	Pid  uint32 `json:"pid"`
	Name string `json:"name"`

	IconUrl       string             `json:"icon_url"`
	Level         uint32             `json:"level"`
	ParentPid     uint32             `json:"parent_pid"`
	Subcategories []PersontagFormattedItem `json:"subs"`
	Parent        *PersontagFormattedItem  `json:"parent"`
}

func FormatPersontag(v m.Persontag) PersontagFormattedItem {
	res := PersontagFormattedItem{
		Name:          v.Name,
		Pid:           publicid.Obfuscate32bit(v.ID),
		Level:         v.Level,
		IconUrl:       v.IconUrl,
		Subcategories: FormatPersontags(v.Sub),
		ParentPid:     publicid.Obfuscate32bit(v.ParentId),
	}
	return res
}

func FormatPersontags(items []m.Persontag) []PersontagFormattedItem {
	res := []PersontagFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPersontag(v))
	}
	return res
}

func FormatPersontagsFromLink(items []m.LinkPersonTag, teamId uint32) []PersontagFormattedItem {
	res := []PersontagFormattedItem{}
	for _, v := range items {
		if v.Tag != nil && v.TeamId == teamId {
			res = append(res, FormatPersontag(*(v.Tag)))
		}
	}
	return res
}
