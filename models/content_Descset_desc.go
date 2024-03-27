package models

import (
	"time"

	"gorm.io/gorm"
)

type LinkDescsetDesc struct {
	ID         uint32         `json:"id" gorm:"primaryKey"`
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
	DescsetId uint32         `json:"descset_id"`
	DescId     uint32         `json:"desc_id"`
	Desc Movedesc `gorm:"foreignKey:desc_id" json:"desc,omitempty"`

}

func (c *LinkDescsetDesc) GetId() uint32 {
	return c.ID
}

func (c *LinkDescsetDesc) TableName() string {
	return "api_content.link_descset_desc"
}
