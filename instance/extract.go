package instance

import (
	"fmt"
)

func (c *CommonInstance) GetConfigFilePath() string {
	return c.GetPathInDevFolder("config/config.yml")
}

func (c *CommonInstance) GetPathInDevFolder(path string) string {
	if c.IsLocalDevEnvironment() {
		return fmt.Sprintf("%v/%v", c.GetDeveloperFolder(), path)

	} else {
		return fmt.Sprintf("/%v", path)
	}
}

func (c *CommonInstance) GetGCPStorageServiceAccountKeysFilePath() string {
	if c.IsLocalDevEnvironment() {
		return fmt.Sprintf("%v/config/gcpstoragekeys.json", c.GetDeveloperFolder())

	} else {
		return "/config/gcpstoragekeys.json"
	}
}

func (c *CommonInstance) GetFilepathOfDevFolrder(filename string) string {
	if c.IsLocalDevEnvironment() {
		return fmt.Sprintf("%v/config/%v", c.GetDeveloperFolder(), filename)

	} else {
		return fmt.Sprintf("/config/%v", filename)

		//		return "/config/gcpstoragekeys.json"
	}
}
