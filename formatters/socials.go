package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type SocialFormattedItem struct {
	Strategy  string `json:"strategy"`
	Pid       uint32 `json:"pid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

func FormatSocial(item m.AuthSocial) SocialFormattedItem {

	return SocialFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		Strategy:  item.Strategy,
		Name:      item.GetName(),
		Email:     item.Email,
		AvatarUrl: item.AvatarUrl,
	}
}

func FormatSocials(items []m.AuthSocial) []SocialFormattedItem {
	res := []SocialFormattedItem{}
	for _, v := range items {
		res = append(res, FormatSocial(v))
	}
	return res
}
