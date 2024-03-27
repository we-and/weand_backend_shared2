package instance

import (
	"fmt"
	"io/ioutil"
	"log"
	config "stretches-common-api/config"

	"gopkg.in/yaml.v2"
)

func (c *CommonInstance) ReadServerConfigFromFile() (config.AppConfig, bool) {
	file := c.GetConfigFilePath()
	fmt.Printf("[%v][config] Config path =%s.\n", c.ServerId, file)
	t := config.AppConfig{}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("ERROR ReadServerConfigFromFile: Cannot read config file path=%v err=%v", file, err)
		return t, false
	}
	err2 := yaml.Unmarshal([]byte(dat), &t)
	if err2 != nil {
		log.Fatalf("ERROR ReadServerConfigFromFile: unmarshalling path=%v %v", file, err2)
		return t, false
	}
	if len(t.Server.JWT_SECRET) == 0 {
		log.Fatalf("ERROR ReadServerConfigFromFile: Invalid config file")
		return t, false
	}
	fmt.Printf("[%v][config] Configuration         : OK\n", c.ServerId)
	return t, true
}
