package formatters

import (
	m "stretches-common-api/models"
)

type DescsetFormattedItem struct {
	Type      string                  `json:"type"`
	Descs []MovedescFormattedItem `json:"descs"`
	Id 	  uint32                  `json:"id"`
}

func FormatDescsets(items []m.Descset, langCode string) []DescsetFormattedItem {
	res := []DescsetFormattedItem{}
	for _, v := range items {
		res = append(res, FormatDescset(v, langCode ))
	}
	return res
}

func FormatDescset(item m.Descset, langCode string) DescsetFormattedItem {
	res := DescsetFormattedItem{

	Id: item.ID,
			Type:  item.Type,
		Descs: FormatDescsFromLink(item.LinkDescs, langCode),
	}
	return res
}
func FormatDescsFromLink(items []m.LinkDescsetDesc, langCode string) []MovedescFormattedItem {
	res := []MovedescFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMovedesc(v.Desc, langCode))
	}
	return res
}
