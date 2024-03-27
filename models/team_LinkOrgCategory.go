package models

type LinkOrgCategory struct {
	ID uint32 `json:"id",gorm:"primaryKey"`

	BatchId    string `json:"batch_id"`
	CategoryId uint32 `json:"category_id"`
	OrgId      uint32 `json:"org_id"`

	//POPULATED
	Category *Category `gorm:"foreignKey:category_id",json:"category,omitempty"`
	Org      *Team     `gorm:"foreignKey:org_id",json:"org,omitempty"`
}

func (c *LinkOrgCategory) GetId() uint32 {
	return c.ID
}

func (c *LinkOrgCategory) TableName() string {
	return "api_team.link_category_org"
}
