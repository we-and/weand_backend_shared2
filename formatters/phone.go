package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type PhoneFormattedItem struct {
	Base        string `json:"base"`
	CountryCode string `json:"country_code"`
	Pid         uint32 `json:"pid"`
}

func FormatPhone(item m.AuthPhone) PhoneFormattedItem {
	return PhoneFormattedItem{
		Pid:         publicid.Obfuscate32bit(item.ID),
		Base:        item.Phone,
		CountryCode: item.CountryCode,
	}
}

func FormatPhones(items []m.AuthPhone) []PhoneFormattedItem {
	res := []PhoneFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPhone(v))
	}
	return res
}
