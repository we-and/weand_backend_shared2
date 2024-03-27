package models

/*
type Org struct {
	ID        uint32         `json:"id",gorm:"primaryKey"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`

	Name string `json:"name"`

	NbInvites uint32 `json:"nb_invites"`
	NbCoaches uint32 `json:"nb_coaches"`
	NbMembers uint32 `json:"nb_members"`

	Invitekey string `json:"invitekey"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	CoverUrl  string `json:"cover_url,omitempty"`

	//POPULATED
	LinkUsers []LinkUserOrg `gorm:"foreignKey:org_id",json:"linkusers,omitempty"`
	LinkTeams []LinkTeamOrg `gorm:"foreignKey:org_id",json:"linkteams,omitempty"`
	Invites   []Invite      `gorm:"foreignKey:org_id",json:"invites,omitempty"`
}

func (c *Org) GetId() uint32 {
	return c.ID
}

func (c *Org) TableName() string {
	return "api_team.org"
}*/
