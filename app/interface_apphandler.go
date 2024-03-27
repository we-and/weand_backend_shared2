package app

import (
	config "stretches-common-api/config"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"

	//"stretches-common-api/querier"
	"gorm.io/gorm"
)

type AppHandlerInterface interface {
	GetDb() *gorm.DB
	GetConfig() *config.AppConfig
	GetFirebase() *firebase.App
	GetGCP() *storage.Client
	//GetQuerier() app.Querier
}
