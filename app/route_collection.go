package app

/*
type RouteHandler struct {
	AppContext *AppContextInterface
}

func CreateRouteHandler(a interface{}) RouteHandler {
	if aa, ok := a.(AppContextInterface); ok {
		return RouteHandler{AppContext: &aa}
	} else {
		return RouteHandler{}
	}
}
func (c *RouteHandler) GetAppContext() AppContextInterface {
	return *c.AppContext
}
func (c *RouteHandler) GetAppContextR() *AppContextInterface {
	return c.AppContext
}
func (c *RouteHandler) GetDb() *gorm.DB {
	if c.AppContext != nil {
		return (*c.AppContext).GetDb()
	} else {
		return nil
	}
}
*/
