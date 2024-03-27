package formatters

import (
	m "stretches-common-api/models"
	"time"
)

type TicketmessageFormattedItem struct {
	CreatedAt    time.Time `json:"created_at"`
	TicketId     uint32    `json:"ticket_id"`
	Content      string    `json:"content"`
	IsRead       bool      `json:"is_read"`
	IsSentByUser bool      `json:"is_sent_by_user"`
}

func FormatTicketmessages(items []m.Ticketmessage) []TicketmessageFormattedItem {
	res := []TicketmessageFormattedItem{}
	for _, v := range items {
		res = append(res, FormatTicketmessage(v))
	}
	return res
}
func FormatTicketmessage(v m.Ticketmessage) TicketmessageFormattedItem {
	res := TicketmessageFormattedItem{
		CreatedAt: v.CreatedAt,
	}

	return res
}
