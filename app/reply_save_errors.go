package app

import (
	"fmt"
	m "stretches-common-api/models"
	"stretches-common-api/response"

	"time"

	"gorm.io/gorm"
)

func SetAndSaveError(r RouteContext, trigger string, err error, errortype string, errCode string) error {
	switch errortype {
	case "INTERNAL_ERROR":
		return SetAndSaveInternalError(r, trigger, err, errCode)
	case "UNAUTHORIZED":
		return SetAndSaveUnauthorized(r, trigger, err, errCode)
	}
	return SetAndSaveInternalError(r, trigger, err, errCode)

}

func SaveError(r RouteContext, trigger string, err error, errCode string) error {
	return SaveErrorGeneric(r, trigger, err, "ERROR", errCode)
}
func SaveUnauthorized(r RouteContext, trigger string, err error, errCode string) error {
	return SaveErrorGeneric(r, trigger, err, "UNAUTHORIZED", errCode)
}
func SaveInternalError(r RouteContext, trigger string, err error, errCode string) error {
	return SaveErrorGeneric(r, trigger, err, "INTERNAL_ERROR", errCode)
}
func SaveBadRequest(r RouteContext, trigger string, err error, errCode string) error {
	return SaveErrorGeneric(r, trigger, err, "BAD_REQUEST", errCode)
}
func SaveWarning(r RouteContext, trigger string, errCode string) error {
	return SaveErrorGeneric(r, trigger, nil, "WARNING", errCode)
}
func SaveNotFound(r RouteContext, trigger string, errCode string) error {
	return SaveErrorGeneric(r, trigger, nil, "NOT_FOUND", errCode)
}
func SaveErrorGeneric(r RouteContext, trigger string, err error, type_ string, errCode string) error {
	config := r.GetConfigR()
	if config == nil {
		return nil
	}
	now := time.Now()
	monitoringError := m.MonitoringError{
		When:    &now,
		Route:   getURL(r.FiberCtx),
		Type:    type_,
		Trigger: trigger,
		Backend: config.BackendId,
		Code:    errCode, //errorcode.Build(r.RouteCode, errCode),
		Desc:    fmt.Sprintf("%v", err),
	}
	db := r.GetDb()
	if db != nil {
		dbres := r.GetDb().Create(&monitoringError)
		if dbres.Error != nil {
			SetInternalError(r, fmt.Sprintf("Error saving error record %v %v", trigger, err), dbres.Error, errCode)
			return nil
		}
	}
	return nil
}

func SetAndSaveWarning(r RouteContext, trigger string, res interface{}, errCode string) error {
	errSave := SaveWarning(r, trigger, errCode)
	OK(r, res)
	return errSave
}
func SetAndSaveNotFound(r RouteContext, db *gorm.DB, backend string, trigger string, err error, errCode string) error {
	SetNotFoundError(r, trigger, err, errCode)
	errSave := SaveNotFound(r, trigger, errCode)
	return errSave
}

func SetAndBadRequest(r RouteContext, db *gorm.DB, backend string, trigger string, err error, errCode string) error {
	SetBadRequest(r, trigger, err, errCode)
	errSave := SaveBadRequest(r, trigger, err, errCode)
	return errSave
}

func SetAndSaveInternalError(r RouteContext, trigger string, err error, errCode string) error {
	SetInternalError(r, trigger, err, errCode)
	errSave := SaveInternalError(r, trigger, err, errCode)
	return errSave
}
func SetAndSaveUnauthorized(r RouteContext, trigger string, err error, errCode string) error {
	response.SetUnauthorized(r.FiberCtx, trigger, err, errCode)
	errSave := SaveUnauthorized(r, trigger, err, errCode)
	return errSave
}

func SetAndSaveBadRequest(r RouteContext, trigger string, err error, errCode string) error {
	response.SetBadRequest(r.FiberCtx, trigger, err, errCode)
	errSave := SaveBadRequest(r, trigger, err, errCode)
	return errSave
}
