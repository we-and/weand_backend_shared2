package formatters

import (
	models "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type ContactFormattedItem struct {
	Pid         uint32                    `json:"pid"`
	FirstName   string                    `json:"first_name,omitempty"`
	LastName    string                    `json:"last_name,omitempty"`
	MiddleName  string                    `json:"middle_name,omitempty"`
	Identifier  string                    `json:"identifier,omitempty"`
	Birthday    string                    `json:"birthday,omitempty"`
	DisplayName string                    `json:"display_name,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
	UpdatedAt   *time.Time                `json:"updated_at,omitempty"`
	DeletedAt   *time.Time                `json:"deleted_at,omitempty"`
	Deleted     bool                      `json:"deleted,omitempty"`
	IsOnMonyl   bool                      `json:"is_on_monyl"`
	UserId      uint32                    `json:"user_id,omitempty"`
	OwnerId     uint32                    `json:"owner_id,omitempty"`
	Details     models.ContactDetailsData `json:"details,omitempty"`
}

func FormatContacts(items []models.Contact) []ContactFormattedItem {
	res := []ContactFormattedItem{}
	for _, v := range items {
		res = append(res, FormatContact(v))
	}
	return res
}
func FormatContact(item models.Contact) ContactFormattedItem {

	return ContactFormattedItem{
		FirstName:  item.FirstName,
		MiddleName: "",
		LastName:   item.LastName,
		Identifier: item.Identifier,

		IsOnMonyl: item.UserId > 0, Details: item.Details,
		Pid: publicid.Obfuscate32bit(item.ID),
	}
}
