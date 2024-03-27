package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type GptJobFormattedItem struct {
	Pid            uint32                 `json:"pid"`
	Id             uint32                 `json:"id"`
	Status         string                 `json:"status"`
	Chats          []GptChatFormattedItem `json:"chats"`
	ElapsedTimeSec int                    `json:"elapsed_time_sec"`
	CreatedAt      *time.Time             `json:"created_at"`
}

func FormatGptJobs(items []m.GPTJob) []GptJobFormattedItem {
	res := []GptJobFormattedItem{}
	for _, v := range items {
		res = append(res, FormatGptJob(v))
	}
	return res
}

func FormatGptJob(item m.GPTJob) GptJobFormattedItem {
	res := GptJobFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		Id:        (item.ID),
		CreatedAt: item.CreatedAt,
		Status:    item.Status,
	}
	if len(item.Chats) > 0 {
		res.Chats = FormatGptChats(item.Chats)
	}
	if item.AIGen != nil {
		res.ElapsedTimeSec = (*(item.AIGen)).ElapsedTimeSec
	}
	return res
}
