package handler

import (
	config "stretches-common-api/config"

	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

type MyHandler struct {
	Db        *gorm.DB
	GCPClient *storage.Client
	Config    config.AppConfig
}

func (h *MyHandler) GetDb() *gorm.DB {
	return h.Db
}
func (h *MyHandler) GetGCP() *storage.Client {
	return h.GCPClient
}
func (h *MyHandler) GetConfig() config.AppConfig {
	return h.Config
}
