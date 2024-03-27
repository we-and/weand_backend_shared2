package log

import (
	"fmt"
	"stretches-common-api/app"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

func LogRequests(fiberapp *fiber.App, h interface{}, dbkey string) {
	if hh, ok := h.(app.AppContextInterface); ok {
		fiberapp.Use(func(c *fiber.Ctx) error {
			r := app.CreateRouteContext(c, hh, "MU256")

			//log request
			searched := fmt.Sprintf("/%s/", dbkey)
			if strings.Contains(c.Path(), searched) {
				go app.SaveLog(r, dbkey)
			}
			return c.Next()
		})
	}
}
