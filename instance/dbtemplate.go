package instance

import (
	"fmt"
	config "stretches-common-api/config"

	"gorm.io/gorm"
)

func (c *CommonInstance) GenerateLocalDB(appConfig config.AppConfig, instances *map[string]*gorm.DB) {
	nick := "local"
	db, localsuccess := c.ConnectToDB(appConfig.Db.Local, nick, false)
	if !localsuccess {
		return
	} else {
		fmt.Printf("[%v][db] Connected to local network.\n", c.ServerId)
	}
	(*instances)[nick] = db
}

func (c *CommonInstance) GenerateLiveDB(appConfig config.AppConfig, instances *map[string]*gorm.DB) {
	nick := "live"
	db, success := c.ConnectToDB(appConfig.Db.Live, nick, false)
	if !success {
		return
	} else {
		fmt.Printf("[%v][db] Connected to live network.\n", c.ServerId)
	}
	(*instances)[nick] = db
}
func (c *CommonInstance) GenerateTestDB(appConfig config.AppConfig, instances *map[string]*gorm.DB) {
	nick := "test"
	db, success := c.ConnectToDB(appConfig.Db.Test, nick, false)
	if !success {
		return
	} else {
		fmt.Printf("[%v][db] Connected to test network.\n", c.ServerId)
	}
	(*instances)[nick] = db
}
