package formatters

import (
	m "stretches-common-api/models"
	"time"

	"stretches-common-api/publicid"
)

type SubscriptionFormattedItem struct {
	CreatedAt *time.Time           `json:"created_at,omitempty"`
	Pid       uint32               `json:"pid"`
	Pricing   PricingFormattedItem `json:"pricing"`
}

func FormatSubscription(v m.Subscription) SubscriptionFormattedItem {

	res := SubscriptionFormattedItem{
		CreatedAt: v.CreatedAt,
		Pid:       publicid.Obfuscate32bit(v.ID),
		Pricing:   FormatPricing(v.Pricing),
	}
	return res
}

func FormatSubscriptions(items []m.Subscription) []SubscriptionFormattedItem {
	res := []SubscriptionFormattedItem{}
	for _, v := range items {
		res = append(res, FormatSubscription(v))
	}
	return res
}
