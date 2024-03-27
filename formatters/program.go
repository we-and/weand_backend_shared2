package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type ProgramFormattedItem struct {
	Name       string                 `json:"name"`
	Desc       string                 `json:"desc"`
	Difficulty string                 `json:"difficulty"`
	Longdesc   string                 `json:"longdesc"`
	Pid        uint32                 `json:"pid"`
	Id         uint32                 `json:"id,omitempty"`
	Workouts   []WorkoutFormattedItem `json:"workouts,omitempty"`
	Plangroups []ChapterFormattedItem `json:"plangroups,omitempty"`
}

func FormatPrograms(items []m.Program, langCode string) []ProgramFormattedItem {
	res := []ProgramFormattedItem{}
	for _, v := range items {
		res = append(res, FormatProgram(v,langCode))
	}
	return res
}
func FormatProgramsForAdmin(items []m.Program, langCode string) []ProgramFormattedItem {
	res := []ProgramFormattedItem{}
	for _, v := range items {
		res = append(res, FormatProgramForAdmin(v,langCode))
	}
	return res
}

/*
	func FormatProgramsFromLinkTeam(links []m.LinkProgramWorkout) []ProgramFormattedItem {
		res := []ProgramFormattedItem{}
		for _, v := range links {
			if v.Program != nil {
				res = append(res, FormatProgram(*v.Program))
			}
		}
		return res
	}
*/
func FormatProgram(item m.Program, langCode string) ProgramFormattedItem {
	res := ProgramFormattedItem{
		Id:         item.ID,
		Name:       item.Name,
		Longdesc:   item.Longdesc,
		Difficulty: item.Difficulty,
		Desc:       item.Desc,
		Pid:        publicid.Obfuscate32bit(item.ID),
	}

	//res.Workouts = FormatWorkoutsFromLinkProgram(item.LinksProgramWorkout)
	res.Plangroups = FormatChapters(item.Plangroups,langCode)
	return res
}

func FormatProgramForAdmin(item m.Program, langCode string) ProgramFormattedItem {
	res := ProgramFormattedItem{
		Name:       item.Name,
		Id:         item.ID,
		Longdesc:   item.Longdesc,
		Difficulty: item.Difficulty,
		Desc:       item.Desc,
		Pid:        publicid.Obfuscate32bit(item.ID),
	}

	//res.Workouts = FormatWorkoutsFromLinkProgram(item.LinksProgramWorkout)
	res.Plangroups = FormatChaptersForAdmin(item.Plangroups,langCode)
	return res
}
