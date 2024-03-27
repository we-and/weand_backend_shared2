package instance

import (
	"errors"
	"log"
	"stretches-common-api/app"
	"strings"

	"gorm.io/gorm"

	fib "github.com/gofiber/fiber/v2"
)

func (cc *CommonInstance) Add404Routes(fiberapp *fib.App, appCtx interface{}) {

	if hh, ok := appCtx.(app.AppContextInterface); ok {
		//	return RouteContext{FiberCtx: c, AppCtx: hh, Db: hh.GetDb(), RouteCode: RouteCode}

		connections := hh.GetDbConnections()
		if co := connections.(*DbConnections); ok {
			instances := co.ConnectionsMap
			//404
			serverId := cc.ServerId
			fiberapp.Use(func(c *fib.Ctx) error {
				r := app.CreateRouteContext(c, hh, "MU201")

				//which database to store the log record on?
				//is url has /test/ then choose test network
				//is url has /live/ then choose live network
				//else choose live network

				if strings.Contains(c.OriginalURL(), "/test/") {
					app.SetAndSaveNotFound(r, (instances)["test"], serverId, "404", errors.New("Not found"), "ME00800")
					return c.SendStatus(404) // => 404 "Not Found"
				} else if strings.Contains(c.OriginalURL(), "/live/") {
					app.SetAndSaveNotFound(r, (instances)["live"], serverId, "404", errors.New("Not found"), "ME00800")
					return c.SendStatus(404) // => 404 "Not Found"
				} else if strings.Contains(c.OriginalURL(), "/local/") {
					app.SetAndSaveNotFound(r, (instances)["local"], serverId, "404", errors.New("Not found"), "ME00800")
					return c.SendStatus(404) // => 404 "Not Found"
				} else {

					var defaultDB *gorm.DB
					if val, ok := (instances)["live"]; ok {
						defaultDB = val
					} else if val, ok := (instances)["test"]; ok {
						defaultDB = val
					} else if val, ok := (instances)["local"]; ok {
						defaultDB = val
					}
					app.SetAndSaveNotFound(r, defaultDB, serverId, "404", errors.New("not found"), "ME00800")
					return c.SendStatus(404) // => 404 "Not Found"
				}
			})
		}
	} else {
		app.CheckAppContext(appCtx)
		log.Fatal("Wrong interface for appContext")
	}
}
