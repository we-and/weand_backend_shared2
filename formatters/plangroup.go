package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type ChapterFormattedItem struct {
	Pid   uint32                `json:"pid"`
	Id    uint32                `json:"id"`
	Idx   uint32                `json:"idx"`
	Name  string                `json:"name"`
	Plans []LessonFormattedItem `json:"plans,omitempty"` //
	// Program ProgramFormattedItem `json:"program,omitempty"`
}

func FormatChapters(items []m.Plangroup, langCode string) []ChapterFormattedItem {
	res := []ChapterFormattedItem{}
	for _, v := range items {
		res = append(res, FormatChapter(v,langCode))
	}
	return res
}
func FormatChaptersForAdmin(items []m.Plangroup, langCode string) []ChapterFormattedItem {
	res := []ChapterFormattedItem{}
	for _, v := range items {
		res = append(res, FormatChapterForAdmin(v,langCode))
	}
	return res
}

func FormatChapter(item m.Plangroup, langCode string) ChapterFormattedItem {

	res := ChapterFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Idx:  item.Idx,
		Name: item.Name,
	}

	//	if(item.Workout!=nil){
	res.Plans = FormatLessons(item.Plans,langCode)
	//	}
	return res
}

func FormatChapterForAdmin(item m.Plangroup, langCode string) ChapterFormattedItem {

	res := ChapterFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Idx:  item.Idx,
		Id:   item.ID,
		Name: item.Name,
	}

	//	if(item.Workout!=nil){
	res.Plans = FormatLessons(item.Plans,langCode)
	//	}
	return res
}
