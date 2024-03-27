package instance

import (
	servers "weand-common-server/servers"

	fib "github.com/gofiber/fiber/v2"
)

func (c *CommonInstance) StartServer(app *fib.App, instances *DbConnections) {
	isHTTPS := c.IsHTTPSProtocol()
	if isHTTPS {
		defer c.CloseConnections(instances.ConnectionsMap)
		servers.ListenHTTPSServer(app, c.ServerAddress, c.ServerId, c.TLSFolder)
	} else {
		c.SetMaxNbConnection(instances.ConnectionsMap)
		defer c.CloseConnections(instances.ConnectionsMap)
		servers.ListenHTTPServer(app, c.ServerAddress, c.ServerId)
	}
}
