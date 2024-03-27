package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	structs "stretches-common-api/structs"
)

type UnregisteredFormattedItem struct {
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name"`
	DisplayName string `json:"display_name"`
	Identifier  string `json:"identifier,omitempty"`
	//	Email       string `json:"email,omitempty"`
	//	Phone       string `json:"phone,omitempty"`
	//	PhoneBase   string `json:"phone_base,omitempty"`
	//	PhoneExt    string `json:"phone_ext,omitempty"`
	//	CountryCode      string                    `json:"countrycode,omitempty"`
	IsYou     bool                `json:"is_you,omitempty"`
	Pid       uint32              `json:"pid"`
	PersonPid uint32              `json:"person_pid"`
//	Details   structs.UserDetails `json:"details,omitempty"`
	//	Details   m.ContactDetailsLightData `json:"details,omitempty"`
}

func FormatUnregistered(u m.Unregistered, me structs.Me, isForEditor bool) UnregisteredFormattedItem {
	///	mar, err := json.Marshal(u.Details)
	res := UnregisteredFormattedItem{
		FirstName:  u.FirstName,
		MiddleName: "",
		LastName:   u.LastName,
		///	Email:       u.Email,
		DisplayName: u.GetDisplayname(),
		//Phone:       u.GetPhone(),
		//	PhoneBase:   u.PhoneBase,
		//	PhoneExt:    u.PhoneExt,
		//Details:     m.ContactDetailsLightData{Emails: []string{u.Email}},
		Pid:   publicid.Obfuscate32bit(u.ID),
		IsYou: me.CheckUserId(u.ID),
	}
	//	if err == nil {
//	res.Details = u.Details //string(mar)
	//	}
	return res
}
