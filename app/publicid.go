package app

import (
	"errors"
	"fmt"
	"strconv"
	"stretches-common-api/publicid"
	"stretches-common-api/validators"

	fiber "github.com/gofiber/fiber/v2"
)

func ParseBody(r RouteContext, req interface{}, code string) (resSuccess bool) {
	if err := r.FiberCtx.BodyParser(req); err != nil {
		SetAndSaveBadRequest(r, "Request fields incorrect.", err, code)
		return
	}
	resSuccess = true
	return
}

func ParseQueryObfuscatedParam(r RouteContext, key string) (bool, uint32) {
	id, errStr, err, errCodePid := Parse32bitObfuscated(r.FiberCtx.Params(key))
	if err != nil {
		SetAndSaveBadRequest(r, errStr, err, errCodePid)
		return false, 0
	}

	return true, id
}

func ParseQueryParam(r RouteContext, key string) (bool, string) {
	return true, r.FiberCtx.Params(key)
}

func ParseIntQueryParam(r RouteContext, key string) (bool, uint32) {
	pidStr := r.FiberCtx.Params(key)
	id64, errConvPid := strconv.ParseUint(pidStr, 10, 64)
	if errConvPid != nil {
		return false, 0
	}
	id32 := uint32(id64)

	return true, id32
}

func Parse32bitObfuscatedPid(c *fiber.Ctx) (uint32, string, error, string) {
	pidStr := c.Params("pid")
	return Parse32bitObfuscated(pidStr)
}

func Parse32bitObfuscatedOrError(r RouteContext, pidStr string) (uint32, bool) { //(id, success) {
	id, errStr, err, errCodePid := Parse32bitObfuscated(r.FiberCtx.Params(pidStr))

	if err != nil {
		SetAndSaveBadRequest(r, errStr, err, errCodePid)
		return 0, false
	}
	return id, true
}
func Parse32bitObfuscated(pidStr string) (uint32, string, error, string) {

	valid, errorDesc := validators.ValidateIdString(pidStr)
	if !valid {
		return 0, fmt.Sprintf("Validation failed for field=%v", errorDesc), errors.New(""), "ME00106"
	}
	pid64, errConvPid := strconv.ParseUint(pidStr, 10, 64)
	if errConvPid != nil {
		return 0, "Cannot turn pidstr in pid=%v : %v", errConvPid, "ME00315"
	}
	pid32 := uint32(pid64)
	id := publicid.Unobfuscate32bit(pid32)
	return id, "", nil, ""
}
