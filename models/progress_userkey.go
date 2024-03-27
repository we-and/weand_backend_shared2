package models

import (
	"time"
)

type UserKey struct {
	ID                   uint32                   `json:"id"`
	UserKey              string                   `json:"user_key,omitempty"`
	UserId               uint32                   `json:"user_id,omitempty"`
	CreatedAt            *time.Time               `json:"created_at,omitempty"`
	UpdatedAt            *time.Time               `json:"updated_at,omitempty"`
	DeletedAt            *time.Time               `json:"deleted_at,omitempty"`
	IsDebug              bool                     `json:"is_debug,omitempty"`
	Devices              []LinkDeviceUser         `gorm:"foreignKey:user_key;references:user_key" json:"linkdevice,omitempty"`
	Profiles             []LinkUserWorkoutprofile `gorm:"foreignKey:user_key;references:user_key" json:"profiles,omitempty"`
	UserAnalytics        UserAnalytics            `gorm:"foreignKey:user_key;references:user_key" json:"useranalytics,omitempty"`
	Notificationsettings NotificationSettings     `json:"notifications_settings,omitempty"`
	TrainingTimes        Trainingtimes            `gorm:"foreignKey:user_key;references:user_key" json:"trainingtimes,omitempty"`
	Iterators            []Iterator               `gorm:"foreignKey:user_key;references:user_key" json:"iterators,omitempty"`
	GptChats             []GPTChat                `gorm:"foreignKey:user_key;references:user_key" json:"gptchats,omitempty"`
	ProgressWorkouts             []ProgressWorkout                `gorm:"foreignKey:user_key;references:user_key" json:"progressworkout,omitempty"`
	FCMTokens             []FirebaseFCMToken                `gorm:"foreignKey:user_key;references:user_key" json:"fcmtokens,omitempty"`
}

func (c *UserKey) GetId() uint32 {
	return c.ID
}
func (c *UserKey) TableName() string {
	return "api_progress.user_key"
}
