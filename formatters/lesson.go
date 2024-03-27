package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type LessonFormattedItem struct {
	Pid          uint32               `json:"pid"`
	Id           uint32               `json:"id"`
	WeekNb       uint32               `json:"week_nb"`
	Dayoftheweek uint32               `json:"dayoftheweek"`
	WorkoutPid   uint32               `json:"workout_pid"`
	WorkoutId    uint32               `json:"workout_id"`
	CoachKey     string               `json:"coach_key"`
	ArenaKey     string               `json:"arena_key"`
	IntroMovePid uint32               `json:"intro_move_pid"`
	Idx          uint32               `json:"idx"`
	IntroMove   MoveFormattedItem    `json:"intro_move,omitempty"`
	ProgramPid   uint32               `json:"program_pid,omitempty"`
	Workout      WorkoutFormattedItem `json:"workout,omitempty"` //
	Tasks        []TaskFormattedItem  `json:"tasks,omitempty"`   //
	// Program ProgramFormattedItem `json:"program,omitempty"`
}

func FormatLessons(items []m.Lesson, langCode string) []LessonFormattedItem {
	res := []LessonFormattedItem{}
	for _, v := range items {
		res = append(res, FormatLesson(v,langCode))
	}
	return res
}

func FormatLesson(item m.Lesson, langCode string) LessonFormattedItem {
	res := LessonFormattedItem{
		Pid:          publicid.Obfuscate32bit(item.ID),
		WeekNb:       item.WeekNb,
		Idx:          item.Idx,
		ArenaKey:     item.ArenaKey,
		Id:           item.ID,
		CoachKey:     item.CoachKey,
		Dayoftheweek: item.Dayoftheweek,
		WorkoutPid:   publicid.Obfuscate32bit(item.WorkoutId),
		WorkoutId:    (item.WorkoutId),
	}
	if item.IntroMoveId>0{
		res.IntroMovePid= publicid.Obfuscate32bit(item.IntroMoveId)
		
	}
		res.IntroMove = FormatMove((item.IntroMove),langCode)
	if item.Workout != nil {
		res.Workout = FormatWorkout(*item.Workout,langCode) 
	}
	res.Tasks = FormatTasks(item.Tasks,langCode)
	return res
}
