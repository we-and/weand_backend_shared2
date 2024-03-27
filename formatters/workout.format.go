package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type WorkoutFormattedItem struct {
	Name        string                   `json:"name"`
	Desc        string                   `json:"desc"`
	Pid         uint32                   `json:"pid"`
	Id          uint32                   `json:"id"`
	ExtraAudios []AudioFormattedItem     `json:"extra_audios,omitempty"`
	CreatedAt   *time.Time               `json:"created_at,omitempty"`
	Movegroups  []MovegroupFormattedItem `json:"movegroups,omitempty"`
	NbMoves     uint32                   `json:"nb_moves"`
	//Durations             map[int]int                   `json:"durations"`
	DurationsMap map[string]interface{} `json:"durations_map"`
	Tags         []TagFormattedItem     `json:"tags,omitempty"`

	// Moves []MoveFormattedItem `json:"moves,omitempty"`
}

func FormatWorkouts(items []m.Workout, langCode string) []WorkoutFormattedItem {
	res := []WorkoutFormattedItem{}
	for _, v := range items {
		res = append(res, FormatWorkout(v, langCode ))
	}
	return res
}
func FormatWorkoutsForAdmin(items []m.Workout, langCode string) []WorkoutFormattedItem {
	res := []WorkoutFormattedItem{}
	for _, v := range items {
		res = append(res, FormatWorkoutForAdmin(v, langCode ))
	}
	return res
}

/*
	func FormatWorkoutsFromLinkProgram(links []m.LinkProgramWorkout) []WorkoutFormattedItem {
		res := []WorkoutFormattedItem{}
		for _, v := range links {
			if v.Workout != nil {
				res = append(res, FormatWorkout(*v.Workout))

			}
		}
		return res
	}
*/
func FormatMovegroupsFromLinkMovegroupWorkout(links []m.LinkMovegroupWorkout, langCode string) []MovegroupFormattedItem {
	res := []MovegroupFormattedItem{}
	for _, v := range links {
		if v.Movegroup != nil {
			f := FormatMovegroup(*v.Movegroup,langCode)
			f.Idx = v.Idx
			res = append(res, f)

		}
	}
	return res
}
func FormatWorkout(item m.Workout, langCode string) WorkoutFormattedItem {
	res := WorkoutFormattedItem{
		Name: item.Name,
		Desc: item.Desc,
		Pid:  publicid.Obfuscate32bit(item.ID),
	}
	res.Tags = FormatTagsFromLinkTagWorkout(item.LinksTagWorkout)
	//res.Durations=item.Durations.Durations
	res.DurationsMap = item.Durations
	res.NbMoves = item.NbMoves
	res.Movegroups = FormatMovegroupsFromLinkMovegroupWorkout(item.LinksMovegroupWorkout, langCode )
	return res
}
func FormatWorkoutForAdmin(item m.Workout, langCode string) WorkoutFormattedItem {
	res := WorkoutFormattedItem{
		CreatedAt: item.CreatedAt,
		Name:      item.Name,
		Desc:      item.Desc,
		Id:        item.ID,
		Pid:       publicid.Obfuscate32bit(item.ID),
	}
	res.Tags = FormatTagsFromLinkTagWorkout(item.LinksTagWorkout)
	//res.Durations=item.Durations.Durations
	res.DurationsMap = item.Durations
	res.NbMoves = item.NbMoves
	res.Movegroups = FormatMovegroupsFromLinkMovegroupWorkout(item.LinksMovegroupWorkout, langCode )
	return res
}
func FormatTagsFromLinkTagWorkout(items []m.LinkTagWorkout) []TagFormattedItem {
	res := []TagFormattedItem{}
	for _, v := range items {
		if v.Tag != nil {
			res = append(res, FormatTag(*v.Tag))
		}
	}
	return res
}
