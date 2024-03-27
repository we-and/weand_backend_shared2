package formatters

import (
	m "stretches-common-api/models"
	"time"
)

type ProfileFormattedItem struct {
	CreatedAt       *time.Time `json:"created_at"`
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	Key             string     `json:"key"`
	DisplaynameMode string     `json:"displayname_mode"`
	AvatarUrl       string     `json:"avatar_url"`
	CoverUrl        string     `json:"cover_url"`
}

func FormatProfile(v m.Profile) ProfileFormattedItem {

	name := ""
	if len(v.Name) > 0 {
		name = v.Name
	} else {
		name = v.User.GetName()
	}
	res := ProfileFormattedItem{
		CreatedAt:       v.CreatedAt,
		Key:             v.Key,
		DisplaynameMode: v.DisplaynameMode,
		//Name: v.User.GetName(),
		AvatarUrl: v.AvatarUrl,
		CoverUrl:  v.CoverUrl,
		Name:      name,
	}
	return res
}
