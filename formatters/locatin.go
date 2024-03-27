package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type LocationFormattedItem struct {
	Long   float64 `json:"long"`
	Lat    float64 `json:"lat"`
	Radius float64 `json:"radius"`
	Pid    uint32  `json:"pid"`
}

func FormatLocations(items []m.Location) []LocationFormattedItem {
	res := []LocationFormattedItem{}
	for _, v := range items {
		res = append(res, FormatLocation(v))
	}
	return res
}
func FormatLocation(item m.Location) LocationFormattedItem {

	return LocationFormattedItem{
		Long:   item.Long,
		Lat:    item.Lat,
		Radius: item.Radius,
		Pid:    publicid.Obfuscate32bit(item.ID),
	}
}
