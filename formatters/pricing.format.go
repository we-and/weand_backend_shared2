package formatters

import (
	"time"

	m "stretches-common-api/models"

	"stretches-common-api/publicid"
)

type PricingFormattedItem struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Pid       uint32     `json:"pid"`
	Price     uint32     `json:"price"`
	Frequency string     `json:"frequency"`
	Currency  string     `json:"currency"`

	Plan PlanFormattedItem `json:"plan"`
}

func FormatPricing(v m.Pricing) PricingFormattedItem {

	res := PricingFormattedItem{
		CreatedAt: v.CreatedAt,
		Price:     v.Price,
		Currency:  v.Currency,
		Frequency: v.Frequency,
		Plan:      FormatPlan(v.Plan),
		Pid:       publicid.Obfuscate32bit(v.ID),
	}
	return res
}

func FormatPricings(items []m.Pricing) []PricingFormattedItem {
	res := []PricingFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPricing(v))
	}
	return res
}
