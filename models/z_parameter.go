package models

type Parameter struct {
	ID    uint32 `json:"id",gorm:"primaryKey"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c *Parameter) GetId() uint32 {
	return c.ID
}

func (c *Parameter) TableName() string {
	return "api_config.parameter"
}
