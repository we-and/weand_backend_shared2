package app

import (
	"fmt"
	m "stretches-common-api/models"

	fiber "github.com/gofiber/fiber/v2"
	//"gorm.io/gorm"
	"time"
)

func GetTokenFromContext(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 {
		return ""
	}
	if authHeader[:7] != "Bearer " {
		return ""
	}
	token := authHeader[7:]
	return token
}
func getDBKey(dbkey string) string {
	if len(dbkey) > 4 {
		return dbkey[0:4]
	} else {
		return dbkey

	}
}
func SaveLog(r RouteContext, dbkey string) error {
	c := r.FiberCtx
	//SAVE LOG
	now := time.Now()
	q := ""
	if len(c.Route().Params) > 0 {
		q = fmt.Sprintf("%v", c.Route().Params)
	}
	log := m.MonitoringLog{
		When: &now,
		//	Route:   c.OriginalURL(),
		IP:      c.IP(),
		URL:     c.Path(),
		Body:    string(c.Request().Body()),
		Network: getDBKey(dbkey),
		Query:   q,
		Service: r.GetConfigR().BackendId,
		Token:   GetTokenFromContext(c),
		Headers: fmt.Sprintf("%v", c.Request().Header.String()),
	}
	{
		dbres := r.GetDb().Create(&log)
		if dbres.Error != nil {
			return SaveError(r, "Cannot save log", dbres.Error, "ME00270")
		}
	}
	return nil
	///	return BadRequest(c, trigger, err)
}
