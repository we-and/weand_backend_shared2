package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
)

type MovegroupFormattedItem struct {
	Name  string              `json:"name"`
	Pid   uint32              `json:"pid"`
	Id    uint32              `json:"id"`
	Idx   uint32              `json:"idx"`
	Moves []MoveFormattedItem `json:"moves,omitempty"`
}

func FormatMovegroups(items []m.Movegroup, langCode string) []MovegroupFormattedItem {
	res := []MovegroupFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMovegroup(v,langCode))
	}
	return res
}

func FormatMovesFromLinkMoveMovegroup(links []m.LinkMovegroupMove, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range links {
		if v.Move != nil {
			it := FormatMove(*v.Move,langCode)
			it.Side = v.Side
			it.Mode = v.Mode
			it.Idx = v.Idx
			///	it.SafetyComment = v.Comment
			//it.SafetyNode = v.SafetyLevel
			res = append(res, it)
		}
	}
	return res
}
func FormatMovegroup(item m.Movegroup, langCode string) MovegroupFormattedItem {

	res := MovegroupFormattedItem{
		Id:   item.ID,
		Name: item.Name,
		Pid:  publicid.Obfuscate32bit(item.ID),
	}
	if len(item.LinksMovegroupMove) > 0 {
		res.Moves = FormatMovesFromLinkMoveMovegroup(item.LinksMovegroupMove,langCode)

	}

	return res
}
