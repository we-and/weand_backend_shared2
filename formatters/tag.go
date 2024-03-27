package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type TagFormattedItem struct {
	//	CreatedAt *time.Time `json:"created_at,omitempty"`
	Pid  uint32 `json:"pid"`
	Name string `json:"name"`
	Type string `json:"type"`
	Key string `json:"key"`
}

func FormatTag(v m.Tag) TagFormattedItem {
	res := TagFormattedItem{
		Type: v.Type,
		Name:          v.Name,
		Pid:           publicid.Obfuscate32bit(v.ID),
		Key:v.Key,
	}
	return res
}

func FormatTags(items []m.Tag) []TagFormattedItem {
	res := []TagFormattedItem{}
	for _, v := range items {
		res = append(res, FormatTag(v))
	}
	return res
}
