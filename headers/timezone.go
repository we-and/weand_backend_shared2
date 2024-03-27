package headers

import (
	"errors"
	"stretches-common-api/app"
	"stretches-common-api/timezone"
	"time"
)

func GetTimezoneNameFromHeader(r app.RouteContext) (string, string, string, error, string) {
	c := r.FiberCtx
	//db := h.GetDb()
	//	config := h.GetConfig()
	//retreive userid
	name := c.Get("X-Timezone-Name")
	if len(name) < 1 {
		return "", "", "Missing timezone name", errors.New(""), "ME00001"
	}
	offset := c.Get("X-Timezone-Offset")
	if len(offset) < 1 {
		return "", "", "Missing timezone name", errors.New(""), "ME00001"
	}
	return name, offset, "", nil, "ME00001"

}
func GetTimezoneOffsetFromHeader(r app.RouteContext) (string, string, string, error, string) {
	c := r.FiberCtx
	//db := h.GetDb()
	//	config := h.GetConfig()
	//retreive userid
	authHeader := c.Get("X-Timezone-Offset")
	if len(authHeader) < 1 {
		return "", "UNAUTHORIZED", "Missing timezone offset", errors.New(""), "ME00001"
	}
	return authHeader, "OK", "", nil, "ME00001"

}

func GetTimezoneOrReturn(r app.RouteContext) (bool, timezone.TzData) {
	tzName, tzOffset, errStr, err, code := GetTimezoneNameFromHeader(r)
	if err != nil {
		app.SetAndSaveUnauthorized(r, errStr, err, code)
		return false, timezone.TzData{}
	}
	loc, err2 := time.LoadLocation(tzName) //ShangHai
	if err2 != nil {
		app.SetAndSaveUnauthorized(r, "Unknown timezone name", err, code)
		return false, timezone.TzData{}
	}
	return true, timezone.TzData{Loc: loc, Name: tzName, Offset: tzOffset}
}
