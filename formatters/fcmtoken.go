package formatters

import m "stretches-common-api/models"


type FCMTokenFormattedItem struct {
	Token string `json:"token"`
	CreatedAt string `json:"created_at"`
}

func FormatFCMTokens(items []m.FirebaseFCMToken) []FCMTokenFormattedItem {
	res := []FCMTokenFormattedItem{}
	for _, v := range items {
		res = append(res, FormatFCMToken(v))
	}
	return res
}
func FormatFCMToken(item m.FirebaseFCMToken) FCMTokenFormattedItem {
	res := FCMTokenFormattedItem{
		Token: item.Token,
		CreatedAt: item.CreatedAt.String(),
	}
	return res
}