package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type LinkMovePropFormattedItem struct {
	Pid    uint32             `json:"pid"`
	Asset  AssetFormattedItem `json:"asset"`
	IsMain bool               `json:"is_main"`
}

func FormatLinkMoveProp(item m.LinkMoveProp,langCode string) LinkMovePropFormattedItem {
	res := LinkMovePropFormattedItem{
		Pid: publicid.Obfuscate32bit(item.ID),
		IsMain: item.IsMain,
		//Character
	}
	if item.Asset != nil {
		res.Asset = FormatAsset(*item.Asset,langCode)
	}
	return res
}

func FormatLinkMoveProps(items []m.LinkMoveProp,langCode string) []LinkMovePropFormattedItem {
	res := []LinkMovePropFormattedItem{}
	for _, v := range items {
		res = append(res, FormatLinkMoveProp(v,langCode))
	}
	return res
}
