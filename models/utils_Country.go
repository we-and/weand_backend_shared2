package models

type Country struct {
	ID     int    `json:"id",gorm:"primaryKey"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Active bool   `json:"active"`
}

func (c *Country) GetId() int {
	return c.ID
}

func (c *Country) TableName() string {
	return "api_utils.country"
}
