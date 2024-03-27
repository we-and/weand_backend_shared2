package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type GoalFormattedItem struct {
	Pid  uint32 `json:"pid"`
	Name string `json:"name"`
	Idx  uint32 `json:"idx"`
	Key  string `json:"key_"`
}

func FormatGoals(items []m.Goal, langCode string) []GoalFormattedItem {
	res := []GoalFormattedItem{}
	for _, v := range items {
		res = append(res, FormatGoal(v, langCode))
	}
	return res
}

func FormatGoal(item m.Goal, langCode string) GoalFormattedItem {
	res := GoalFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Idx:  item.Idx,
		Name: item.GetName(langCode),
		Key:  item.Key,
	}
	return res
}
