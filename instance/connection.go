package instance

import (
	"fmt"
	"log"
	"os"
	"time"

	config "stretches-common-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func (c *CommonInstance) ConnectToDB(instanceDb config.DatabaseInstanceConfig, nick string, useSSL bool) (*gorm.DB, bool) {
	fmt.Printf("[%v][db] Connecting to DB %v \n", c.ServerId, nick)

	dsn := ""

	if useSSL {
		//prepareThreshold=0 for no cache. otherwise any change in schema would trigger a cache invalidation error
		//		dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=require sslrootcert=%v sslkey=%v sslcert=%v application_name=%s  prepareThreshold=0",
		dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=require sslrootcert=%v sslkey=%v sslcert=%v application_name=%s TimeZone=UTC prepareThreshold=0 	",
			instanceDb.POSTGRES_USER,
			instanceDb.POSTGRES_PASSWORD,
			instanceDb.POSTGRES_HOST,
			instanceDb.POSTGRES_DB,
			instanceDb.POSTGRES_PORT,
			c.getServerCaPath(instanceDb.CertificateFolder),
			c.getClientKeyPath(instanceDb.CertificateFolder),
			c.getClientCertPath(instanceDb.CertificateFolder),
			c.ServerLongId,
		)
	} else {
		dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d  TimeZone=UTC",
			instanceDb.POSTGRES_USER,
			instanceDb.POSTGRES_PASSWORD,
			instanceDb.POSTGRES_HOST,
			instanceDb.POSTGRES_DB,
			instanceDb.POSTGRES_PORT)
	}
	//	fmt.Printf("[%v] Connecting to DB %v with \n\t%s\n", c.ServerId, nick, strings.Replace(dsn, " ", "\n\t", -1))
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			//	IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful: false, // Disable color
		},
	)
	db, dberr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			//		currentTime := time.Now()
			//		_, offset := currentTime.Zone()
			//		mysqlTime := currentTime.Add(time.Second * time.Duration(offset))
			now := time.Now().UTC()
			return now
		},
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "api.", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,   // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if dberr != nil {
		return db, false
	}
	fmt.Printf("[%v][db] Connected to DB %v     : OK\n", c.ServerId, nick)
	return db, true
}

func (c *CommonInstance) SetMaxNbConnection(instances map[string]*gorm.DB) {
	isLocal := c.IsLocalDevEnvironment()
	if isLocal {
		for key, db := range instances {
			fmt.Printf("[%v][db] Nb connections for %s: 2\n", c.ServerId, key)
			sqlDb, sqlErr := db.DB()
			if sqlErr != nil {
				sqlDb.SetMaxIdleConns(1)
				sqlDb.SetMaxOpenConns(1)
			}
		}
	} else {
		fmt.Printf("[%v][db] Nb connections: AUTO\n", c.ServerId)
	}

}

func (c *CommonInstance) CloseConnections(instances map[string]*gorm.DB) {

	fmt.Printf("[%v][db] Close connection\n", c.ServerId)

	for dbkey, db := range instances {
		sqlDb, sqlErr := db.DB()
		if sqlErr != nil {
			//			sqlDb.SetMaxIdleConns()
			fmt.Printf("[%v][db] Closing connection %s\v", c.ServerId, dbkey)
			defer sqlDb.Close()
		}
	}
}
