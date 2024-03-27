package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Movedesc struct {
	ID         uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
	Complexity int            `json:"complexity"`
	MoveId     uint32         `json:"move_id"`
	HasAudio   bool           `json:"has_audio"`
	AudioList  string         `json:"audio_list"`
	Content    string         `json:"content"`
	Title      string         `json:"title"`
	Idx        int            `json:"idx"`
	Audios []Audio `gorm:"foreignKey:desc_id" json:"audios,omitempty"`
	Translations []Translation `gorm:"foreignKey:obj_id"`
}

func (c *Movedesc) GetContent(langCode string) string {
	if langCode != "" {
		for _, k := range c.Translations {
			if k.Type == "MOVE_DESC_TEXT" {
				return k.Text
			}
		}
	}
	return c.Content

}
func (c *Movedesc) GetId() uint32 {
	return c.ID
}

func (c *Movedesc) TableName() string {
	return "api_content.movedesc"
}
func (c *Movedesc) GetAudioList() []string {
	return strings.Split(c.AudioList, ",")
}


func (c *Movedesc) AddToAudioList(name string) {
	if strings.Contains(c.AudioList, name) {
		return
	}
	if c.AudioList == "" {
		c.AudioList = name
	} else {
		c.AudioList = c.AudioList + "," + name
	}
}
