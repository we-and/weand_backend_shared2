package formatters

import (
	m "stretches-common-api/models"

	"stretches-common-api/publicid"
)

type TaskFormattedItem struct {
	Pid       uint32               `json:"pid"`
	Idx       uint32               `json:"idx"`
	LessonPid uint32               `json:"lesson_pid"`
	ItemPid   uint32               `json:"item_pid"`
	Points    int                  `json:"points"`
	Type      string               `json:"type"`
	Title     string               `json:"title"`
	Workout   WorkoutFormattedItem `json:"workout"`
	Article   ArticleFormattedItem `json:"article"`
}

func FormatTasks(items []m.Task, langCode string) []TaskFormattedItem {
	res := []TaskFormattedItem{}
	for _, v := range items {
		res = append(res, FormatTask(v, langCode))
	}
	return res
}
func FormatTask(v m.Task, langCode string) TaskFormattedItem {
	res := TaskFormattedItem{
		Pid:       publicid.Obfuscate32bit(v.ID),
		LessonPid: publicid.Obfuscate32bit(v.LessonId),
		ItemPid:   publicid.Obfuscate32bit(v.ItemId),
		Points:    v.Points,
		Type:      v.Type,
		Title:     v.Title,
		Idx:       v.Idx,
	}
	if v.ItemWorkout != nil {
		res.Workout = FormatWorkout(*v.ItemWorkout, langCode)
	}
	if v.ItemReading != nil {
		res.Article = FormatArticle(*v.ItemReading, langCode)
	}
	return res
}
