package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type ExperienceFormattedItem struct {
	Pid  uint32 `json:"pid"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Idx  uint32 `json:"idx"`
	Key  string `json:"key_"`
}

func FormatExperiences(items []m.Experience, langCode string) []ExperienceFormattedItem {
	res := []ExperienceFormattedItem{}
	for _, v := range items {
		res = append(res, FormatExperience(v, langCode))
	}
	return res
}

func FormatExperience(item m.Experience, langCode string) ExperienceFormattedItem {
	res := ExperienceFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Idx:  item.Idx,
		Name: item.GetName(langCode),
		Desc: item.GetDesc(langCode),
		Key:  item.Key,
	}
	return res
}
