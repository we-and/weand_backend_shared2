package app

import (
	config "stretches-common-api/config"

	//	"stretches-common-api/querier"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RouteContext struct {
	FiberCtx  *fiber.Ctx
	AppCtx    AppContextInterface
	Db        *gorm.DB
	RouteCode string
	Querier   interface{}
}

func CreateRouteContext(c *fiber.Ctx, h interface{}, RouteCode string) RouteContext {
	if hh, ok := h.(AppContextInterface); ok {
		return RouteContext{FiberCtx: c, AppCtx: hh, Db: hh.GetDb(), RouteCode: RouteCode}
	} else {
		CheckAppContext(h)
		return RouteContext{FiberCtx: c, RouteCode: RouteCode}
	}
}

func CreateRouteContextWithQuerier(c *fiber.Ctx, h interface{}, RouteCode string, querier interface{}) RouteContext {
	if hh, ok := h.(AppContextInterface); ok {
		return RouteContext{FiberCtx: c, AppCtx: hh, Db: hh.GetDb(), RouteCode: RouteCode, Querier: querier}
	} else {
		return RouteContext{FiberCtx: c, RouteCode: RouteCode}
	}
}
func (r *RouteContext) GetConfigR() *config.AppConfig {
	if r.AppCtx == nil {
		return nil
	}
	return ((r.AppCtx).GetConfig())
}

func (r *RouteContext) GetConfigA() config.AppConfig {
	return (*(r.AppCtx).GetConfig())
}

func (r *RouteContext) GetDb() *gorm.DB {
	if r.AppCtx == nil {
		return nil
	}
	return (r.AppCtx).GetDb()
}
func (r *RouteContext) AppContext() AppContextInterface {
	return (r.AppCtx)
}
func (r *RouteContext) Firebase() *firebase.App {
	if r.AppCtx == nil {
		return nil
	}
	return (r.AppCtx).GetFirebase()
}

func (r *RouteContext) AppInstance() interface{} {
	if r.AppCtx == nil {
		return nil
	}
	return (r.AppCtx).GetAppInstance()
}
