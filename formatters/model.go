package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type ModelFormattedItem struct {
	Pid      uint32              `json:"pid"`
	Filename string              `json:"filename"`
	Bucket   string              `json:"bucket"`
	Url      string              `json:"url,omitempty"`
	AxisUp   string              `json:"axis_up,omitempty"`
	AssetPid uint32              `json:"asset_pid"`
	Asset    *AssetFormattedItem `json:"asset,omitempty"`
	AnimKey  string              `json:"anim_key,omitempty"`
}

func FormatModel(item m.Model,langCode string) ModelFormattedItem {
	res := ModelFormattedItem{
		//CharacterPid:  publicid.Obfuscate32bit(item.CharacterId),
		Pid:      publicid.Obfuscate32bit(item.ID),
		AnimKey:  item.AnimKey,
		Filename: item.Filename,
		Url:      item.Url,
		AxisUp:   item.AxisUp,
		Bucket:   item.Bucket,
	}
	if item.Asset != nil && item.Asset.ID > 0 {
		f := FormatAsset(*item.Asset,langCode )
		res.Asset = &f
	}
	return res
}
func FormatModels(items []m.Model,langCode string) []ModelFormattedItem {
	res := []ModelFormattedItem{}
	for _, v := range items {
		res = append(res, FormatModel(v,langCode ))
	}
	return res
}
