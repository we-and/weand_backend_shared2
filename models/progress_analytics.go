package models

import (
	_ "time"

	"gorm.io/gorm"
)

type UserAnalytics struct {
	gorm.Model
	ID                  uint32 `json:"id" gorm:"primaryKey"`
	UserId              uint32 `json:"user_id"`
	UserKey             string `json:"user_key"`
	NbWorkoutsStarted   uint32 `json:"nb_workouts_started"`
	SecondsTrained      int    `json:"seconds_trained"`
	NbWorkoutsEnded     uint32 `json:"nb_workouts_ended"`
	NbMovesTrained      uint32 `json:"nb_moves_trained"`
	NbMovesCompleted    uint32 `json:"nb_moves_completed"`
	NbMovesStarted      uint32 `json:"nb_moves_started"`
	NbLogins            uint32 `json:"nb_logins"`
	NbSessions          uint32 `json:"nb_sessions"`
	NbGenerateStarted   uint32 `json:"nb_generate_started"`
	NbGenerateCompleted uint32 `json:"nb_generate_completed"`
	NbInstalls          uint32 `json:"nb_installs"`
}

func (c *UserAnalytics) GetId() uint32 {
	return c.ID
}

func (c *UserAnalytics) TableName() string {
	return "api_progress.user_analytics"
}
