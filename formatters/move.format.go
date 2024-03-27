package formatters

import (
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"strings"
	"unicode"
)

type AudioFormattedItem struct {
	Filename   string `json:"filename"`
	Voice      string `json:"voice"`
	Country    string `json:"country"`
	Id         uint32 `json:"id"`
	Pid        uint32 `json:"pid"`
	Type       string `json:"type"`
	DurationMs int    `json:"duration_ms"`
}

func FormatAudios(items []m.Audio) []AudioFormattedItem {
	res := []AudioFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAudio(v))
	}
	return res
}
func FormatAudio(item m.Audio) AudioFormattedItem {
	res := AudioFormattedItem{
		DurationMs: item.DurationMs,
		Filename:   item.Filename,
		Voice:      item.Voice,
		Country:    item.Country,
		Id:         item.ID,
		Pid:        publicid.Obfuscate32bit(item.ID),
		Type:       item.Type,
	}
	return res
}

type MovedescFormattedItem struct {
	HasAudio  bool                 `json:"has_audio"`
	Content   string               `json:"content"`
	AudioList []string             `json:"audio_list"`
	Audios    []AudioFormattedItem `json:"audios"`
	Title     string               `json:"title"`
	Idx       int                  `json:"idx"`
	Id        uint32               `json:"id"`
	Pid       uint32               `json:"pid"`
}

type MovecommentFormattedItem struct {
	SafetyLevel string `json:"safety_level"`
	Comment     string `json:"comment"`
}

func FormatMovedescs(items []m.Movedesc, langCode string) []MovedescFormattedItem {
	res := []MovedescFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMovedesc(v, langCode))
	}
	return res
}
func FormatMovecomments(items []m.LinkMoveProfileraw) []MovecommentFormattedItem {
	res := []MovecommentFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMovecomment(v))
	}
	return res
}

func FormatMovedesc(item m.Movedesc, langCode string) MovedescFormattedItem {
	res := MovedescFormattedItem{
		Pid:       publicid.Obfuscate32bit(item.ID),
		Id:        item.ID,
		AudioList: item.GetAudioList(),
		Content:   item.GetContent(langCode),
		Audios:    FormatAudios(item.Audios),
		Title:     item.Title,
		HasAudio:  item.HasAudio,
		Idx:       item.Idx,
	}
	return res
}

func FormatMovecomment(item m.LinkMoveProfileraw) MovecommentFormattedItem {
	res := MovecommentFormattedItem{
		SafetyLevel: item.SafetyLevel,
		Comment:     item.Comment,
	}
	return res
}

type VariantGroupFormattedItem struct {
	Members []MoveFormattedItem `json:"members,omitempty"`
	Id      uint32              `json:"id,omitempty"`
}
type MoveFormattedItem struct {
	Name          string                      `json:"name"`
	Name2         string                      `json:"name2"`
	Name3         string                      `json:"name3"`
	Name4         string                      `json:"name4"`
	Name5         string                      `json:"name5"`
	Name6         string                      `json:"name6"`
	Name7         string                      `json:"name7"`
	SearchNames   string                      `json:"search_names"`
	AudioList     string                      `json:"audio_list"`
	Level         uint32                      `json:"level,omitempty"`
	ViewCount     uint32                      `json:"view_count"`
	ModelCount    int                         `json:"model_count"`
	NameAudios    []AudioFormattedItem        `json:"name_audios,omitempty"`
	WorkoutCount  int                         `json:"workout_count"`
	ImageUrl      string                      `json:"image_url,omitempty"`
	DescCount     int                         `json:"desc_count"`
	AudioCount    int                         `json:"audio_count"`
	Desc          string                      `json:"desc,omitempty"`
	Translations  []TranslationFormattedItem  `json:"translations,omitempty"`
	Descs         []MovedescFormattedItem     `json:"descs,omitempty"`
	Desc2         string                      `json:"desc2,omitempty"`
	AnimKey       string                      `json:"anim_key,omitempty"`
	Position      string                      `json:"position,omitempty"`
	Mistakes      string                      `json:"mistakes,omitempty"`
	Benefits      string                      `json:"benefits,omitempty"`
	Variations    string                      `json:"variations,omitempty"`
	IsVariant     bool                        `json:"is_variant,omitempty"`
	Warning       string                      `json:"warning,omitempty"`
	Tips          string                      `json:"tips,omitempty"`
	SanskritName  string                      `json:"sanskrit_name,omitempty"`
	IsChiral      bool                        `json:"is_chiral,omitempty"`
	Variants      []MoveFormattedItem         `json:"variants,omitempty"`
	Physios       []PhysioFormattedItem       `json:"physios,omitempty"`
	VariantGroups []VariantGroupFormattedItem `json:"variantgroups,omitempty"`
	Props         []LinkMovePropFormattedItem `json:"props,omitempty"`
	Models        []ModelFormattedItem        `json:"models,omitempty"`
	Movecomments  []MovecommentFormattedItem  `json:"movecomments,omitempty"`
	Pid           uint32                      `json:"pid"`
	Idx           uint32                      `json:"idx,omitempty"`
	Id            uint32                      `json:"id"`
	Descsets      []DescsetFormattedItem      `json:"descsets,omitempty"`
	Side          string                      `json:"side,omitempty"`
	Mode          string                      `json:"mode,omitempty"`
	// SafetyLevel    string                      `json:"safety_node,omitempty"`
	// SComment string                      `json:"safety_comment,omitempty"`
}
type TranslationFormattedItem struct {
	Language string `json:"language"`
	ObjId    uint32 `json:"obj_id"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

func FormatTranslations(t []m.Translation) []TranslationFormattedItem {
	res := []TranslationFormattedItem{}
	for _, v := range t {
		res = append(res, FormatTranslation(v))

	}
	return res
}
func FormatTranslation(t m.Translation) TranslationFormattedItem {
	res := TranslationFormattedItem{
		Language: t.Lang,
		ObjId:    t.ObjId,
		Type:     t.Type,
		Content:  t.Text,
	}
	return res
}
func FormatMoves(items []m.Move, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMove(v, langCode))
	}
	return res
}
func FormatMovesForAdmin(items []m.Move, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMoveForAdmin(v, langCode))
	}
	return res
}
func FormatMovesSimple(items []m.Move, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range items {
		res = append(res, FormatMoveSimple(v, langCode))
	}
	return res
}

func FormatMovesFromLinkWorkout(links []m.LinkMoveWorkout, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range links {
		if v.Move != nil {
			it := FormatMove(*v.Move, langCode)
			it.Idx = v.Idx

			res = append(res, it)
		}
	}
	return res
}
func FormatMove(item m.Move, langCode string) MoveFormattedItem {
	res := MoveFormattedItem{
		Name:         item.GetName(langCode),
		Name2:        item.Name2,
		Name3:        item.Name3,
		DescCount:    item.DescCount,
		ImageUrl:     item.ImageUrl,
		Variations:   item.Variations,
		Warning:      item.Warning,
		Desc:         item.Desc,
		Desc2:        item.Desc2,
		Position:     item.Position,
		AudioList:    item.AudioList,
		Id:           item.ID,
		Benefits:     item.Benefits,
		Mistakes:     item.Mistakes,
		Tips:         item.Tips,
		SanskritName: item.SanskritName,
		IsChiral:     item.IsChiral,
		AnimKey:      item.AnimKey,

		Pid: publicid.Obfuscate32bit(item.ID),
	}

	//	if len(item.Descs) > 0 {
	//		res.Descs = FormatMovedescs(item.Descs, langCode)
	//	}
	if len(item.NameAudios) > 0 {
		res.NameAudios = FormatAudios(item.NameAudios)
	}
	if len(item.Descsets) > 0 {
		res.Descsets = FormatDescsets(item.Descsets, langCode)
	}
	if len(item.LinksProfile) > 0 {
		res.Movecomments = FormatMovecomments(item.LinksProfile)
	}
	if len(item.Variants) > 0 {
		res.Variants = FormatMoves(item.Variants, langCode)
	}

	if len(item.Physios) > 0 {
		res.Physios = FormatPhysios(item.Physios,langCode)
	}
	if len(item.Models) > 0 {
		res.Models = FormatModels(item.Models, langCode)
	}
	if len(item.Props) > 0 {
		res.Props = FormatLinkMoveProps(item.Props, langCode)
	}

	return res
}
func FormatMoveForAdmin(item m.Move, langCode string) MoveFormattedItem {
	res := MoveFormattedItem{
		DescCount:    item.DescCount,
		AudioCount:   item.AudioCount,
		WorkoutCount: item.WorkoutCount,
		ModelCount:   item.ModelCount,
		Name:         item.Name,
		Name2:        item.Name2,
		Name3:        item.Name3,
		Name4:        item.Name4,
		Name5:        item.Name5,
		Name6:        item.Name6,
		Name7:        item.Name7,
		ImageUrl:     item.ImageUrl,
		Variations:   item.Variations,
		Warning:      item.Warning,
		Desc:         item.Desc,
		Desc2:        item.Desc2,
		IsVariant:    len(item.VariantMembers) > 0,
		Position:     item.Position,
		AudioList:    item.AudioList,
		Id:           item.ID,
		ViewCount:    item.NbViews,
		SearchNames:  item.SearchNames,
		Benefits:     item.Benefits,
		Mistakes:     item.Mistakes,
		Tips:         item.Tips,
		SanskritName: item.SanskritName,
		IsChiral:     item.IsChiral,
		AnimKey:      item.AnimKey,

		Pid: publicid.Obfuscate32bit(item.ID),
	}
	res.Translations = FormatTranslations(item.Translations)
	if len(item.Descs) > 0 {
		res.Descs = FormatMovedescs(item.Descs, langCode)
	}
	//	res.Props = FormatLinkMoveProps(item.Props)
	if len(item.LinksProfile) > 0 {
		res.Movecomments = FormatMovecomments(item.LinksProfile)
	}
	if len(item.Variants) > 0 {
		res.Variants = FormatMoves(item.Variants, langCode)
	}
	if len(item.NameAudios) > 0 {
		res.NameAudios = FormatAudios(item.NameAudios)
	}
	if len(item.VariantMembers) > 0 {
		res.VariantGroups = FormatVariantGroups(item.VariantMembers, langCode)
	}
	if len(item.Physios) > 0 {
		res.Physios = FormatPhysios(item.Physios,langCode)
	}
	if len(item.Models) > 0 {
		res.Models = FormatModels(item.Models, langCode)
	}
	if len(item.Props) > 0 {
		res.Props = FormatLinkMoveProps(item.Props, langCode)
	}
	if len(item.Descsets) > 0 {
		res.Descsets = FormatDescsets(item.Descsets, langCode)
	}

	return res
}
func FormatVariantGroups(variantMembers []m.VariantMember, langCode string) []VariantGroupFormattedItem {
	res := []VariantGroupFormattedItem{}
	for _, v := range variantMembers {
		res = append(res, FormatVariantGroupFromMember(v, langCode))
	}
	return res
}
func FormatVariantGroupFromMember(variantMember m.VariantMember, langCode string) VariantGroupFormattedItem {
	res := VariantGroupFormattedItem{}
	if variantMember.Group != nil {
		res.Id = variantMember.Group.ID
		res.Members = FormatVariantMembersFromGroup(*variantMember.Group, langCode)
	}
	return res
}
func FormatVariantMembersFromGroup(variantGroup m.VariantGroup, langCode string) []MoveFormattedItem {
	res := []MoveFormattedItem{}
	for _, v := range variantGroup.GroupMembers {
		//	if v.Move != nil {
		res = append(res, FormatMoveForAdmin(v.RelatedMove, langCode))
		//	}
	}
	return res
}

func FormatMoveSimple(item m.Move, langCode string) MoveFormattedItem {
	res := MoveFormattedItem{
		Name:     item.Name,
		Id:       item.ID,
		ImageUrl: item.ImageUrl,
		IsChiral: item.IsChiral,
		AnimKey:  item.AnimKey,
		Pid:      publicid.Obfuscate32bit(item.ID),
	}

	return res
}

func MoveGetSearchName(name string) string {
	res := strings.ReplaceAll(name, " ", "")
	res = MoveAlphanumericOnly(res)
	return res
}
func MoveAlphanumericOnly(str string) string {
	var result []rune
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return strings.ToLower(string(result))
}
