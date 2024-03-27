package models

import (
	"gorm.io/gorm"
)

type TimeToRender struct {
	gorm.Model

	ID uint32 `json:"id" gorm:"primaryKey"`

	DeviceId       string `json:"device_id,omitempty"`
	InstanceId     uint32 `json:"instance_id,omitempty"`
	DownloadTimeMs int    `json:"download_time_ms,omitempty"`
	ModelId        int    `json:"model_id,omitempty"`
	RenderTimeMs   int    `json:"render_time_ms,omitempty"`
	LoadingTimeMs  int    `json:"loading_time_ms,omitempty"`
}

func (c *TimeToRender) GetId() uint32 {
	return c.ID
}

func (c *TimeToRender) TableName() string {
	return "api_analytics.time_to_render"
}
