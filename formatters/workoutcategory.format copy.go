package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type WorkoutCategoryFormattedItem struct {
	Name     string                         `json:"name"`
	IconUrl  string                         `json:"icon_url"`
	Pid      uint32                         `json:"pid"`
	Level    uint32                         `json:"level"`
	Desc     string                         `json:"desc"`
	Workouts []WorkoutFormattedItem         `json:"workouts,omitempty"`
	Children []WorkoutCategoryFormattedItem `json:"children,omitempty"`
}

func FormatWorkoutCategories(items []m.WorkoutCategory, langCode string) []WorkoutCategoryFormattedItem {
	res := []WorkoutCategoryFormattedItem{}
	for _, v := range items {
		res = append(res, FormatWorkoutCategory(v, langCode ))
	}
	return res
}
func FormatWorkoutsFromLinkCategoryWorkout(links []m.LinkCategoryWorkout, langCode string) []WorkoutFormattedItem {
	res := []WorkoutFormattedItem{}
	for _, v := range links {
		if v.Workout != nil {
			res = append(res, FormatWorkout(*v.Workout, langCode ))

		}
	}
	return res
}
func FormatWorkoutCategory(item m.WorkoutCategory, langCode string) WorkoutCategoryFormattedItem {

	res := WorkoutCategoryFormattedItem{
		Name:    item.Name,
		Desc:    item.Desc,
		Level:   item.Level,
		IconUrl: item.IconUrl,
		Pid:     publicid.Obfuscate32bit(item.ID),
	}
	res.Workouts = FormatWorkoutsFromLinkCategoryWorkout(item.LinksCategoryWorkout, langCode )
	res.Children = FormatWorkoutCategories(item.Children, langCode)
	return res
}
