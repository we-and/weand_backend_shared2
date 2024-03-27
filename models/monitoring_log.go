package models

import "time"

type MonitoringLog struct {
	ID   int        `json:"id" gorm:"primaryKey"`
	When *time.Time `json:"when,omitempty"`
	URL  string     `json:"url,omitempty"`
	IP   string     `json:"ip,omitempty"`
	//	Route string  `json:"route,omitempty"`
	Body    string `json:"body,omitempty"`
	Query   string `json:"query,omitempty"`
	Headers string `json:"headers,omitempty"`
	Token   string `json:"token,omitempty"`
	Network string `json:"network,omitempty"`
	Service string `json:"service,omitempty"`
}

func (c *MonitoringLog) GetId() int {
	return c.ID
}

func (c *MonitoringLog) TableName() string {
	return "api_monitoring.logs"
}
