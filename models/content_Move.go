package models

import (
	"strings"
	"time"
	"unicode"

	"gorm.io/gorm"
)

type Move struct {
	ID        uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	IsGenerated  bool   `json:"is_generated"`
	AudioList    string `json:"audio_list"`
	Name         string `json:"name"`
	WorkoutCount int    `json:"workout_count"`
	ModelCount   int    `json:"model_count"`
	Desc         string `json:"desc"`
	DescCount    int    `json:"desc_count"`
	AudioCount   int    `json:"audio_count"`
	Desc2        string `json:"desc2"`
	IsChiral     bool   `json:"is_chiral"`
	ParentId     uint32 `json:"parent_id"`
	Name2        string `json:"name2"`
	ImageUrl     string `json:"image_url"`
	Name3        string `json:"name3"`
	Name4        string `json:"name4"`
	Name5        string `json:"name5"`
	Name6        string `json:"name4"`
	Name7        string `json:"name5"`
	SearchNames  string `json:"search_names"`
	Variations   string `json:"variations"`

	Benefits     string `json:"benefits,omitempty"`
	Mistakes     string `json:"mistakes,omitempty"`
	Purpose      string `json:"purpose,omitempty"`
	Progression  string `json:"progression,omitempty"`
	Warning      string `json:"warning,omitempty"`
	Tips         string `json:"tips,omitempty"`
	SanskritName string `json:"sanskrit_name,omitempty"`

	Level    uint32 `json:"level"`
	AnimKey  string `json:"anim_key"`
	NbViews  uint32 `json:"nb_views"`
	Position string `json:"position"`

	Variants []Move `gorm:"foreignKey:parent_id" json:"variants,omitempty"`
	//MoveVariants []MoveVariant  `gorm:"foreignKey:move_id" json:"movevariants,omitempty"`
	VariantMembers []VariantMember `gorm:"foreignKey:move_id" json:"variantmembers,omitempty"`
	//VariantMember []VariantMember `gorm:"foreignKey:move_id" json:"variantmember,omitempty"`
	Physios    []Physio       `gorm:"foreignKey:move_id" json:"physios,omitempty"`
	Props      []LinkMoveProp `gorm:"foreignKey:move_id" json:"props,omitempty"`
	Descs      []Movedesc     `gorm:"foreignKey:move_id" json:"movedescs,omitempty"`
	Models     []Model        `gorm:"foreignKey:move_id" json:"models,omitempty"`
	NameAudios []Audio        `gorm:"foreignKey:move_id" json:"nameaudios,omitempty"`

	Translations []Translation        `gorm:"foreignKey:move_id"`
	LinksProfile []LinkMoveProfileraw `gorm:"foreignKey:move_id" json:"linkprofile,omitempty"`
	Descsets     []Descset            `gorm:"foreignKey:move_id" json:"descset,omitempty"`
	// Movedescs []Movedesc `gorm:"foreignKey:move_id" json:"movedescs,omitempty"`
}

func (c *Move) GetId() uint32 {
	return c.ID
}

func (c *Move) GetName(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "MOVE_NAME" {
				return k.Text
			}
		}
	}
	return c.Name

}

func (c *Move) ResetNames() {
	// searchnames := []string{}
	c.SearchNames = ""
	names := []string{
		c.Name,
		c.Name2,
		c.Name3,
		c.Name4,
		c.Name5,
		c.Name6,
		c.Name7,
		c.SanskritName,
	}
	for _, name := range names {
		c.AddSearchName(name)
	}
}
func (c *Move) SetNameAsAdditional(name string) string {
	if name == "" {
		return ""
	}
	if c.Name2 == "" {
		c.Name2 = name
		return "name2"
	}
	if c.Name3 == "" {
		c.Name3 = name
		return "name3"
	}
	if c.Name4 == "" {
		c.Name4 = name
		return "name4"
	}
	if c.Name5 == "" {
		c.Name5 = name
		return "name5"
	}
	if c.Name6 == "" {
		c.Name6 = name
		return "name6"
	}
	if c.Name7 == "" {
		c.Name7 = name
		return "name7"
	}
	return ""
}
func (c *Move) AddSearchName(name string) {
	if name == "" {
		return
	}
	newsearchname := GetSearchName(name)
	if c.SearchNames == "" {
		c.SearchNames = newsearchname
	} else {

		//check if already exists
		searchnames := strings.Split(c.SearchNames, "||")
		for _, v := range searchnames {
			if v == newsearchname {
				return
			}
		}
		c.SearchNames = c.SearchNames + "||" + newsearchname
	}
}

func (c *Move) TableName() string {
	return "api_content.move"
}

func GetSearchName(name string) string {
	res := strings.ReplaceAll(name, " ", "")
	res = AlphanumericOnly(res)
	return res
}
func AlphanumericOnly(str string) string {
	var result []rune
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return strings.ToLower(string(result))
}
