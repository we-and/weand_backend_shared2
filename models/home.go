package models

type Home struct {
	ID    int    `json:"id",gorm:"primaryKey"`
	Key   string `json:"key",omitempty`
	Fr_CA string `json:"fr_ca",omitempty`
	En_CA string `json:"en_ca",omitempty`
}

func (c *Home) GetId() int {
	return c.ID
}

func (c *Home) TableName() string {
	return "api_translations.home"
}
