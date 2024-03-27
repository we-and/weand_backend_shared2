package formatters

import (
	m "stretches-common-api/models"

	"stretches-common-api/publicid"
)
type PhysioFormattedItem struct {
	Pid      uint32                 `json:"pid"`
	Anatomy []AnatomyFormattedItem `json:"anatomy"`
	Desc string `json:"desc"`
	Idx uint32 `json:"idx"`
}

func FormatPhysio(item m.Physio, langCode string) PhysioFormattedItem {
	res := PhysioFormattedItem{
		Pid:  publicid.Obfuscate32bit(item.ID),
		Desc:item.Desc,
		Idx:item.Idx,
	}
	anat:=[]AnatomyFormattedItem{}
	for _, k:= range  item.LinksAnatomy {
		if k.Anatomy != nil {
			anat=append(anat, FormatAnatomy(*k.Anatomy,langCode))
		}
	}
	res.Anatomy =anat

	return res
}


func FormatPhysios(items []m.Physio, langCode string) []PhysioFormattedItem {
	res := []PhysioFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPhysio(v, langCode ))
	}
	return res
}
