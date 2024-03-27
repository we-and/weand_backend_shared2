package app

import (
	"fmt"
	config "stretches-common-api/config"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"gorm.io/gorm"
)

type AppContextInterface interface {
	GetDb() *gorm.DB
	GetConfig() *config.AppConfig
	GetPublicConfig() *config.AppConfig
	GetFirebase() *firebase.App
	GetGCP() *storage.Client
	GetAppInstance() interface{}
	GetDbConnections() interface{}
}

// ////////////////////////////////////////////////////////////////
// DEBUG
// ////////////////////////////////////////////////////////////////
type AppContextInterfaceA interface {
}

type AppContextInterfaceDb interface {
	GetDb() *gorm.DB
}
type AppContextInterfaceConfig interface {
	GetConfig() *config.AppConfig
}
type AppContextInterfaceGCP interface {
	GetGCP() *storage.Client
}
type AppContextInterfaceFirebase interface {
	GetFirebase() *firebase.App
}
type AppContextInterfacePublicConfig interface {
	GetPublicConfig() *config.AppConfig
}
type AppContextInterfaceAppInstance interface {
	GetAppInstance() interface{}
}
type AppContextInterfaceDbConnections interface {
	GetDbConnections() interface{}
}

func CheckAppContext(a interface{}) {
	if _, ok := a.(AppContextInterfaceDb); ok {
		fmt.Println("AppContextInterfaceDb ok")
	}
	if _, ok := a.(AppContextInterfaceConfig); ok {
		fmt.Println("AppContextInterfaceConfig ok")
	}
	if _, ok := a.(AppContextInterfacePublicConfig); ok {
		fmt.Println("AppContextInterfacePublicConfig ok")
	}
	if _, ok := a.(AppContextInterfaceGCP); ok {
		fmt.Println("AppContextInterfaceGCP ok")
	}
	if _, ok := a.(AppContextInterfaceFirebase); ok {
		fmt.Println("AppContextInterfaceFirebase ok")
	}
	if _, ok := a.(AppContextInterfaceAppInstance); ok {
		fmt.Println("AppContextInterfaceAppInstance ok")
	}
	if _, ok := a.(AppContextInterfaceDbConnections); ok {
		fmt.Println("AppContextInterfaceDbConnections ok")
	}
	if _, ok := a.(AppContextInterface); ok {
		fmt.Print("AppContextInterface ok")
	}
}
