package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type GptChatFormattedItem struct {
	Pid       uint32     `json:"pid"`
	Id        uint32     `json:"id"`
	JobId     uint32     `json:"job_id"`
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
}

func FormatGptChats(items []m.GPTChat) []GptChatFormattedItem {
	res := []GptChatFormattedItem{}
	for _, v := range items {
		res = append(res, FormatGptChat(v))
	}
	return res
}

func FormatGptChat(item m.GPTChat) GptChatFormattedItem {
	res := GptChatFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		Id:        (item.ID),
		Role:      item.Role,
		JobId:     item.JobId,
		Content:   item.Content,
		CreatedAt: item.CreatedAt,
	}
	return res
}
