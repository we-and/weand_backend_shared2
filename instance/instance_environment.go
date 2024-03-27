package instance

import (
	"fmt"
	"os"
	config "stretches-common-api/config"

	common "stretches-common-api"
	utils "weand-common-server/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (c *CommonInstance) GenerateDBConnections() DbConnections {
	connections := map[string]*gorm.DB{}
	appConfig := c.AppConfig
	c.GenerateLiveDB(appConfig, &connections)
	c.GenerateTestDB(appConfig, &connections)
	//	if c.IsLocalDevEnvironment() {
	//		c.GenerateLocalDB(appConfig, &connections)
	//	}
	dbConnections := DbConnections{
		ConnectionsMap: connections,
	}
	c.DbConnections = &dbConnections
	return dbConnections
}
func (c *CommonInstance) GenerateLiveAndTestDBConnections() DbConnections {
	connections := map[string]*gorm.DB{}
	appConfig := c.AppConfig
	c.GenerateLiveDB(appConfig, &connections)
	c.GenerateTestDB(appConfig, &connections)
	//	if c.IsLocalDevEnvironment() {
	//		c.GenerateLocalDB(appConfig, &connections)
	//	}
	dbConnections := DbConnections{
		ConnectionsMap: connections,
	}
	c.DbConnections = &dbConnections
	return dbConnections
}
func (c *CommonInstance) WelcomeAndInit() *config.AppConfig {
	fmt.Printf("[%v] API server init\n", c.ServerId)
	c.PrintEnvVars()
	///Read yaml config file
	configpath := c.GetConfigFilePath()
	fmt.Printf("[%v] Config path =%s.\n", c.ServerId, configpath)

	appConfig, success2 := c.ReadServerConfigFromFile()
	if !success2 {
		return nil
	}

	appConfig.ConfigFolder = c.GetDeveloperFolder()
	appConfig.BackendId = c.ServerLongId
	c.AppConfig = appConfig
	return &appConfig
}
func (c *CommonInstance) PrintEnvVars() {
	arr := []string{
		c.ProjConfig.DEV_FOLDER_KEY,
		c.ProjConfig.ENV_KEY,
		c.ProjConfig.PROTOCOL_KEY,
	}

	c.PrintEnvVarsWithKey(arr)
}
func (c *CommonInstance) IsLocalDevEnvironment() bool {
	return c.IsLocalDevEnvironmentWithKey(c.ProjConfig.ENV_KEY)
}
func (c *CommonInstance) IsHTTPSProtocol() bool {
	return c.CheckEnvVar(c.ProjConfig.PROTOCOL_KEY, "HTTPS")
}
func (c *CommonInstance) GenerateServer(limitRate bool, printRoutes bool) *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: printRoutes,
	})

	///Server specific config
	utils.SetHeaders(app)
	if limitRate {
		utils.ApplyRateLimiter(app)
	}
	utils.ApplyLogger(app)
	///API Server specific config
	common.ApplyDocumentation(app, fmt.Sprintf("https://api.weand.co.uk/%v", c.ServerId))
	return app
}
func (c *CommonInstance) GetDeveloperFolder() string {
	DEV_FOLDER := os.Getenv(c.ProjConfig.DEV_FOLDER_KEY)
	return DEV_FOLDER
}
