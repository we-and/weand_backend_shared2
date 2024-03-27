package instance

import "gorm.io/gorm"

type DbConnections struct {
	ConnectionsMap map[string]*gorm.DB
}
