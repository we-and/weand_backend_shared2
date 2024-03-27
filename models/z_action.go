package models

type Action struct {
	ID   uint32 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (c *Action) GetId() uint32 {
	return c.ID
}

func (c *Action) TableName() string {
	return "api_content.action"
}
