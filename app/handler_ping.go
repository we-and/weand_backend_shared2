package app

import (
	"fmt"
	"stretches-common-api/types"

	response "stretches-common-api/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddPingRoute(app *fiber.App, serverId string, version string, dbKey string, db *gorm.DB) {

	type DbTimezoneResult struct {
		TimeZone string
	}

	var result DbTimezoneResult
	db.Raw("show timezone").Scan(&result)

	pingResponse := types.PingResponse{
		Timezone: result.TimeZone,
		Version:  version, Network: dbKey, Service: serverId, Status: "Live"}
	v := fmt.Sprintf("[%v][route] Add ping on /%v/%v\n", serverId, serverId, dbKey)
	fmt.Print(v)
	app.Get(fmt.Sprintf("/%v/%v", serverId, dbKey), func(c *fiber.Ctx) error {

		return c.JSON(pingResponse)
	})

}

type serverReconnectResponse struct {
	Success bool `json:"success"`
}

func AddReconnectRoute(app *fiber.App, serverId string, version string, dbKey string, db *gorm.DB) {
	app.Get("/reconnect", func(c *fiber.Ctx) error {
		response.SetOK(c, serverReconnectResponse{Success: true})
		return nil
	})
}
