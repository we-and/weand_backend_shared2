package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type AssetFormattedItem struct {
	Pid          uint32               `json:"pid"`
	Tags         []TagFormattedItem   `json:"tags"`
	Name         string               `json:"name"`
	IsGear       bool                 `json:"is_gear"`
	Key          string               `json:"key"`
	Lighting     string               `json:"lighting,omitempty"`
	ThumbnailUrl string               `json:"thumbnail_url"`
	Models       []ModelFormattedItem `json:"models"`
}

func FormatAsset(item m.Asset, langCode string) AssetFormattedItem {
	res := AssetFormattedItem{
		Pid:          publicid.Obfuscate32bit(item.ID),
		Name:         item.GetName(langCode),
		IsGear:       item.IsGear,
		ThumbnailUrl: item.ThumbnailUrl,
		Key:          item.Key,
		//Character
	}
	res.Lighting = item.Lighting
	res.Models = FormatModels(item.LinksModel,langCode)
	res.Tags = FormatTagsFromLinks(item.LinksTag)
	return res
}

func FormatTagsFromLinks(items []m.LinkAssetTag) []TagFormattedItem {
	res := []TagFormattedItem{}
	for _, v := range items {
		if v.Tag != nil {
			res = append(res, FormatTag(*v.Tag))
		}
	}
	return res
}
func FormatAssets(items []m.Asset, langCode string) []AssetFormattedItem {
	res := []AssetFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAsset(v, langCode))
	}
	return res
}

func FormatAssetsFromLinks(items []m.LinkMoveProp, langCode string) []AssetFormattedItem {
	res := []AssetFormattedItem{}
	for _, v := range items {
		if v.Asset != nil {
			res = append(res, FormatAsset(*v.Asset, langCode))
		}
	}
	return res
}
