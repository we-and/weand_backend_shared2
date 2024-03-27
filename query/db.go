package query

import (
	"fmt"
	"stretches-common-api/app"
	"stretches-common-api/errorcode"

	"reflect"

	"gorm.io/gorm"
)

func Create(r app.RouteContext, db *gorm.DB, obj interface{}, queryCode string) bool { //success
	result := db.Create(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot create %v", name)
		//FORMAT ERROR OUTPUT
		errCode := errorcode.Build(r.RouteCode, queryCode)
		app.SetAndSaveInternalError(r, msg, result.Error, errCode)
		return false
	}
	return true
}
func CreateOrRollback(r app.RouteContext, db *gorm.DB, obj interface{}, code string) bool { //success
	result := db.Create(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot create %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func Save(r app.RouteContext, db *gorm.DB, obj interface{}, code string) bool { //success
	result := db.Save(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot save %v", name)
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func SaveOrRollback(r app.RouteContext, db *gorm.DB, obj interface{}, code string) bool { //success
	result := db.Save(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot save %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback after save tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func Delete(r app.RouteContext, db *gorm.DB, obj interface{}, queryCode string) bool { //success
	result := db.Delete(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot delete %v", name)
		//FORMAT ERROR OUTPUT
		errCode := errorcode.Build(r.RouteCode, queryCode)

		app.SetAndSaveInternalError(r, msg, result.Error, errCode)
		return false
	}
	return true
}

func DeleteOrRollback(r app.RouteContext, db *gorm.DB, obj interface{}, code string) bool { //success
	result := db.Delete(obj)
	name := reflect.TypeOf(obj).Elem().Name()
	msg := fmt.Sprintf("Cannot delete %v", name)
	if result.Error != nil {
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback after delete tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func FirstWhere(r app.RouteContext, db *gorm.DB, obj interface{}, queryCode string) bool { //success
	result := db.First(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot fetch %v", name)
		//FORMAT ERROR OUTPUT
		errCode := errorcode.Build(r.RouteCode, queryCode)
		app.SetAndSaveInternalError(r, msg, result.Error, errCode)
		return false
	}
	return true
}
func FirstWhereOrRollback(r app.RouteContext, db *gorm.DB, obj interface{}, code string) bool { //success
	result := db.First(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot fetch %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func FindWhere(r app.RouteContext, db *gorm.DB, obj interface{}, queryCode string) bool { //success
	result := db.Find(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot find %v", name)
		//FORMAT ERROR OUTPUT
		errCode := errorcode.Build(r.RouteCode, queryCode)
		app.SetAndSaveInternalError(r, msg, result.Error, errCode)
		return false
	}
	return true
}
func FindWhereOrRollback(r app.RouteContext, db *gorm.DB, obj []interface{}, code string) bool { //success
	result := db.Find(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot find %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func CountWhere(r app.RouteContext, db *gorm.DB, obj *int64, queryCode string) bool { //success
	result := db.Count(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot count %v", name)
		//FORMAT ERROR OUTPUT
		errCode := errorcode.Build(r.RouteCode, queryCode)
		app.SetAndSaveInternalError(r, msg, result.Error, errCode)
		return false
	}
	return true
}

func CountWhereOrRollback(r app.RouteContext, db *gorm.DB, obj *int64, code string) bool { //success
	result := db.Count(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot count %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback tx", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func UpdateWhere(r app.RouteContext, db *gorm.DB, obj *map[string]interface{}, code string) bool { //success
	result := db.Updates(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot update %v", name)
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func UpdatesWhereOrRollback(r app.RouteContext, db *gorm.DB, obj *int64, code string) bool { //success
	result := db.Updates(obj)
	if result.Error != nil {
		name := reflect.TypeOf(obj).Elem().Name()
		msg := fmt.Sprintf("Cannot update %v", name)
		if errRoll := db.Rollback().Error; errRoll != nil {
			app.SetAndSaveInternalError(r, "rollback tx after update", errRoll, fmt.Sprintf("%v-rollback", code))
			return false
		}
		//FORMAT ERROR OUTPUT
		app.SetAndSaveInternalError(r, msg, result.Error, code)
		return false
	}
	return true
}

func ParsePost(r app.RouteContext, obj interface{}, code string) bool {
	if err := (*r.FiberCtx).BodyParser(&obj); err != nil {
		app.SetAndSaveBadRequest(r, "Incorrect POST request fields.", err, code)
		return false
	}
	return true
}
