package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type IteratorFormattedItem struct {
	CreatedAt    *time.Time           `json:"created_at"`
	PlangroupPid uint32               `json:"plangroup_pid"`
	ProgramPid   uint32               `json:"program_pid"`
	Pid          uint32               `json:"pid"`
	WorkoutPid   uint32               `json:"workout_pid"`
	PlangroupId  uint32               `json:"plangroup_id"`
	ProgramId    uint32               `json:"program_id"`
	WorkoutId    uint32               `json:"workout_id"`
	Program      ProgramFormattedItem `json:"program,omitempty"`
}

func FormatIterator(v m.Iterator, langCode string) IteratorFormattedItem {
	res := IteratorFormattedItem{}
	res = IteratorFormattedItem{
		CreatedAt:    v.CreatedAt,
		Pid:          publicid.Obfuscate32bit(v.ID),
		PlangroupPid: publicid.Obfuscate32bit(v.PlangroupId),
		ProgramPid:   publicid.Obfuscate32bit(v.ProgramId),
		WorkoutPid:   publicid.Obfuscate32bit(v.WorkoutId),
	}
	if v.Program != nil {
		res.Program = FormatProgramForAdmin(*v.Program,langCode)
	}
	return res
}
func FormatIteratorForAdmin(v m.Iterator, langCode string) IteratorFormattedItem {
	res := IteratorFormattedItem{}
	res = IteratorFormattedItem{
		Pid:         publicid.Obfuscate32bit(v.ID),
		CreatedAt:   v.CreatedAt,
		PlangroupId: (v.PlangroupId),
		ProgramId:   (v.ProgramId),
		WorkoutId:   (v.WorkoutId),
	}
	if v.Program != nil {
		res.Program = FormatProgramForAdmin(*v.Program,langCode)
	}
	return res
}
func FormatIteratorsForAdmin(items []m.Iterator, langCode string) []IteratorFormattedItem {
	res := []IteratorFormattedItem{}
	for _, v := range items {
		res = append(res, FormatIteratorForAdmin(v,langCode))
	}
	return res
}
