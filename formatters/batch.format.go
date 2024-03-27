package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
	timezone "stretches-common-api/timezone"
	"time"
)

type BatchFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	//	DeletedAt *time.Time         `json:"deleted_at,omitempty"`
	//	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	Pid  uint32             `json:"pid"`
	Jobs []JobFormattedItem `json:"jobs,omitempty"`
}

func FormatBatches(items []m.SendBatch, me structs.Me, reqTzData *timezone.TzData) []BatchFormattedItem {
	res := []BatchFormattedItem{}
	for _, v := range items {
		res = append(res, FormatBatch(v, me, reqTzData))
	}
	return res
}
func FormatBatch(item m.SendBatch, me structs.Me, reqTzData *timezone.TzData) BatchFormattedItem {

	res := BatchFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		CreatedAt: item.CreatedAt,
		//		UpdatedAt: item.UpdatedAt,
		//		DeletedAt: item.DeletedAt,
		Jobs: FormatJobs(item.Jobs, me, reqTzData),
	}

	return res
}
