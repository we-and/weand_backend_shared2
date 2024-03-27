package instance

import (
	"fmt"
	"io/ioutil"
)

func (c *CommonInstance) getClientCertPath(folder string) string {
	isLocalDev := c.IsLocalDevEnvironment()
	if isLocalDev {
		return c.getClientCertPathLocal(folder)
	} else {
		return fmt.Sprintf("/tls/%s/clientcert.pem", folder)
	}
}

func (c *CommonInstance) getClientKeyPath(folder string) string {
	isLocalDev := c.IsLocalDevEnvironment()
	if isLocalDev {
		return c.getClientKeyPathLocal(folder)
	} else {
		return fmt.Sprintf("/tls/%s/clientkey.pem", folder)
	}
}
func (c *CommonInstance) getServerCaPath(folder string) string {
	isLocalDev := c.IsLocalDevEnvironment()
	if isLocalDev {
		return c.getServerCaPathLocal(folder)
	} else {
		return fmt.Sprintf("/tls/%s/serverca.pem", folder)
	}
}
func (c *CommonInstance) checkFileExists(path string) bool {
	_, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot read file at %v", path))
	}
	return true
}
func (c *CommonInstance) getClientCertPathLocal(folder string) string {
	MONYL_DEV_FOLDER := c.GetDeveloperFolder()
	path := fmt.Sprintf("%v/%s/ssl/client-cert.pem", MONYL_DEV_FOLDER, folder)
	c.checkFileExists(path)
	fmt.Print("SSL Client cert       : OK\n")
	return path
}
func (c *CommonInstance) getClientKeyPathLocal(folder string) string {
	DEV_FOLDER := c.GetDeveloperFolder()
	path := fmt.Sprintf("%v/%s/ssl/client-key.pem", DEV_FOLDER, folder)
	c.checkFileExists(path)
	fmt.Print("SSL Client key        : OK\n")
	return path
}
func (c *CommonInstance) getServerCaPathLocal(folder string) string {
	DEV_FOLDER := c.GetDeveloperFolder()
	path := fmt.Sprintf("%v/%s/ssl/server-ca.pem", DEV_FOLDER, folder)
	c.checkFileExists(path)
	fmt.Print("SSL Server cert       : OK\n")
	return path
}
