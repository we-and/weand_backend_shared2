package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type AnatomyfocusFormattedItem struct {
	Pid      uint32 `json:"pid"`
	IsCustom bool   `json:"is_custom"`
	Idx      uint32 `json:"idx"`
	Name     string `json:"name"`
}
type AnatomyFormattedItem struct {
	Pid       uint32                      `json:"pid"`
	Name      string                      `json:"name"`
	Idx       uint32                      `json:"idx"`
	Level     uint32                      `json:"level"`
	Posx      float32                     `json:"posx"`
	Posy      float32                     `json:"posy"`
	Labelx    float32                     `json:"labelx"`
	Labely    float32                     `json:"labely"`
	ParentPid uint32                      `json:"parent_pid"`
	Focus     []AnatomyfocusFormattedItem `json:"focus"`
}

func FormatAnatomies(items []m.Anatomy, langCode string) []AnatomyFormattedItem {
	res := []AnatomyFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAnatomy(v, langCode))
	}
	return res
}
func FormatAnatomyfocuses(items []m.Anatomyfocus) []AnatomyfocusFormattedItem {
	res := []AnatomyfocusFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAnatomyfocus(v))
	}
	return res
}

func FormatAnatomy(item m.Anatomy, langCode string) AnatomyFormattedItem {

	res := AnatomyFormattedItem{
		Pid:    publicid.Obfuscate32bit(item.ID),
		Idx:    item.Idx,
		Name:   item.GetName(langCode),
		Posx:   item.Posx,
		Posy:   item.Posy,
		Labelx: item.Labelx,
		Labely: item.Labely,
		Focus:  FormatAnatomyfocuses(item.AnatomyFocus),
	}
	return res
}

func FormatAnatomyfocus(item m.Anatomyfocus) AnatomyfocusFormattedItem {

	res := AnatomyfocusFormattedItem{
		Pid:      publicid.Obfuscate32bit(item.ID),
		Idx:      item.Idx,
		IsCustom: item.IsCustom,
		Name:     item.Name,
	}
	return res
}
