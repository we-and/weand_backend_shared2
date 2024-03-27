package instance

import (
	"fmt"
	"os"
)

func (c *CommonInstance) PrintEnvVarsWithKey(keys []string) {
	for _, v := range keys {
		fmt.Printf("[%v][ENV] %v =\t%s\n", c.ServerId, v, os.Getenv(v))
	}
}
func (c *CommonInstance) IsLocalDevEnvironmentWithKey(env string) bool {
	KEY_ENV := os.Getenv(env)
	return KEY_ENV == "LOCAL"
}
func (c *CommonInstance) CheckEnvVar(env string, key string) bool {
	KEY_ENV := os.Getenv(env)
	return KEY_ENV == key
}

func (c *CommonInstance) IsHTTPSWithKey(env string) bool {
	KEY_ENV := os.Getenv(env)
	return KEY_ENV == "HTTPS"
}
func (c *CommonInstance) GetDeveloperFolderWithKey(devfolder string) string {
	DEV_FOLDER := os.Getenv(devfolder)
	return DEV_FOLDER
}
