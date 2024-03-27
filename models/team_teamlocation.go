package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamLocation struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`

	Name    string  `json:"name"`
	Address string  `json:"address"`
	PlaceId string  `json:"place_id"`
	TeamId  uint32  `json:"team_id"`
	Long    float64 `json:"long"`
	Lat     float64 `json:"lat"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (c *TeamLocation) GetId() uint32 {
	return c.ID
}

func (c *TeamLocation) TableName() string {
	return "api_team.teamlocation"
}
