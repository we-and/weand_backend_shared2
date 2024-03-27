package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type CategoryFormattedItem struct {
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
	Pid     uint32 `json:"pid"`

	Level         uint32                  `json:"level"`
	ParentPid     uint32                  `json:"parent_pid"`
	Subcategories []CategoryFormattedItem `json:"subcategories"`
	Parent        *CategoryFormattedItem  `json:"parent"`
}

func FormatCategories(items []m.Category) []CategoryFormattedItem {
	res := []CategoryFormattedItem{}
	for _, v := range items {
		res = append(res, FormatCategory(v))
	}
	return res
}

func FormatCategoriesFromLinkTeam(links []m.LinkTeamCategory) []CategoryFormattedItem {
	res := []CategoryFormattedItem{}
	for _, v := range links {
		if v.Category != nil {
			res = append(res, FormatCategory(*v.Category))

		}
	}
	return res
}
func FormatCategory(item m.Category) CategoryFormattedItem {

	res := CategoryFormattedItem{
		Name:          item.Name,
		Level:         item.Level,
		IconUrl:       item.IconUrl,
		Subcategories: FormatCategories(item.Sub),
		ParentPid:     publicid.Obfuscate32bit(item.ParentId),
		Pid:           publicid.Obfuscate32bit(item.ID),
	}
	if item.Parent != nil {
		p := FormatCategory(*item.Parent)
		res.Parent = &p
	}
	return res
}
