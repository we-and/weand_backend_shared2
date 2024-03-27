package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type SportFormattedItem struct {
	Pid       uint32 `json:"pid"`
	Name      string `json:"name"`
	Idx       uint32 `json:"idx"`
	Level     int    `json:"level"`
	ParentPid uint32 `json:"parent_pid"`
	Key       string `json:"key_"`
}

func FormatSports(items []m.Sport, langCode string) []SportFormattedItem {
	res := []SportFormattedItem{}
	for _, v := range items {
		res = append(res, FormatSport(v, langCode))
	}
	return res
}

func FormatSport(item m.Sport, langCode string) SportFormattedItem {
	res := SportFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		Idx:       item.Idx,
		Level:     item.Level,
		ParentPid: publicid.Obfuscate32bit(item.ParentId),

		Name: item.GetName(langCode),
		Key:  item.Key,
	}
	return res
}
