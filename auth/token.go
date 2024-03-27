package auth

import (
	"errors"
	"fmt"
	"regexp"
	"stretches-common-api/app"
	config "stretches-common-api/config"
	m "stretches-common-api/models"
	"stretches-common-api/query"
	"stretches-common-api/structs"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// /RETURNS (jwt, success)
func GenerateJWT(upublicId string, config config.AppConfig) (string, bool) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["publicid"] = upublicId
	claims["exp"] = time.Now().Add(time.Hour * 7200).Unix() //30 days

	t, errValidation := token.SignedString([]byte(config.Server.JWT_SECRET))
	if errValidation != nil {
		return "", false
	}
	return t, true

}

func fetchUserIdFromPublicId(db *gorm.DB, pid string) (uint32, string, error) { //(userId, errorMsg,error)
	var count int64
	{
		result := db.Model(&m.PublicId{}).Where("publickey = ?", pid).Count(&count)
		if result.Error != nil {
			return 0, fmt.Sprintf("Cannot count publicid where publickey=%v %v", pid, result.Error), result.Error
		}
	}
	if count > 0 {
		publicid := m.PublicId{}
		{
			result := db.Where("publickey = ?", pid).First(&publicid)
			if result.Error != nil {
				return 0, fmt.Sprintf("Cannot fetch publicid where publickey=%v. PublicId expired? %v", pid, result.Error), result.Error
			}
		}
		return publicid.UserID, "", nil
	} else {
		return 0, fmt.Sprintf("No publicid for publickey=%v  PublicId expired or wrong network?", pid), errors.New("")
	}
}
func GetUserIdFromToken(db *gorm.DB, secret string, tokenString string) (resSuccess bool, resUserId uint32, resErrStr string, resErr error) {
	pid, resErrStr, resErr := getUserPublicIdFromToken(secret, tokenString)
	if resErr != nil {
		return
	}
	id, resErrStr, resErr := fetchUserIdFromPublicId(db, pid)
	if resErr != nil {
		return
	}
	resSuccess = true
	resUserId = id
	return
}
func getUserPublicIdFromToken(secret string, tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", "Error parsing auth token", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		publicidinterface := claims["publicid"]
		publicIdStr := fmt.Sprintf("%v", publicidinterface)
		return publicIdStr, "", nil
	} else {
		return "", "Invalid token", errors.New("Invalid token")
	}
}

func AllowConfirmedUsers(r app.RouteContext) (bool, structs.Me) {
	hasAccess, me, errStr, errAuth := ConfirmedUserAuthWall(r)
	if !hasAccess {
		app.SetAndSaveUnauthorized(r, errStr, errAuth, "MX000-001")
		return false, me
	}
	return hasAccess, me
}

func AllowConfirmedUsersOrUserKey(r app.RouteContext) (bool, structs.Me) {
	hasAccess, me, errStr, errAuth := ConfirmedUserAuthWall(r)
	hasKey, userKey := GetUserKeyFromHeader(r)
	if hasKey {
		me.UserKey = userKey
	}
	if !hasAccess {
		if hasKey {
			db := r.GetDb()
			p := m.Person{}
			if !query.FirstWhere(r, db.Where("user_key = ?", userKey), &p, "ME006-010") {
				p.UserKey = userKey
				if !query.Create(r, db, &p, "ME006-010") {
					return false, me //, "INTERNAL_ERROR", "Could not create person", errors.New(""), "ME006-010"
				}
				return false, me //, "INTERNAL_ERROR", "No person with user_id", errors.New(""), "ME006-010"
			}
			return hasKey, structs.Me{UserKey: userKey, PersonId: p.ID}
		} else {
			app.SetAndSaveUnauthorized(r, errStr, errAuth, "MX000-001")
			return false, me

		}
	}
	return hasAccess, me
}

func AllowAdminUsers(r app.RouteContext) (bool, structs.Me) {
	hasAccess, me, errStr, errAuth := AdminAuthWall(r)
	if !hasAccess {
		app.SetAndSaveUnauthorized(r, errStr, errAuth, "MX000-001")
		return false, me
	}
	return hasAccess, me
}
func ConfirmedUserAuthWall(r app.RouteContext) (bool, structs.Me, string, error) {
	//auth wall
	success, me, errtype, errstr, err, errCode := GetUserIdAndRightsFromHeader(r)
	if !success {
		app.SaveErrorGeneric(r, errstr, err, errtype, errCode)
		return false, me, errstr, err
	}

	//check rights
	if !HasUserRights(me.Rights) {
		str := "Access forbidden. Confirmed users only."
		app.SaveUnauthorized(r, str, errors.New(""), "ME00025")
		return false, me, str, errors.New("")
	}
	return true, me, "", nil

}
func AdminAuthWall(r app.RouteContext) (bool, structs.Me, string, error) {
	//get rights
	success, me, errtype, errstr, err, errCode := GetUserIdAndRightsFromHeader(r)
	if !success {
		app.SaveErrorGeneric(r, errstr, err, errtype, errCode)
		return false, me, errstr, err
	}

	//check rights
	if !HasAdminRights(me.Rights) {
		str := "Access forbidden. Admin users only."
		app.SaveUnauthorized(r, str, errors.New(""), "ME00026")
		return false, me, str, errors.New("")
	}
	return true, me, "", nil
}

// (userId , Rights, status, errDesc, error, code)
func GetUserKeyFromHeader(r app.RouteContext) (bool, string) {
	c := r.FiberCtx
	userKeyHeader := c.Get("UserKey")
	if len(userKeyHeader) < 5 {
		return false, "ME00001"
	}
	return true, userKeyHeader
}

// (userId , Rights, status, errDesc, error, code)
func GetUserIdAndRightsFromHeader(r app.RouteContext) (bool, structs.Me, string, string, error, string) {
	me := structs.Me{}
	c := r.FiberCtx
	db := r.GetDb()
	if db == nil {
		return false, me, "INTERNAL_ERROR", "No db", errors.New(""), "ME00001"
	}
	config := r.GetConfigR()
	if config == nil {
		return false, me, "INTERNAL_ERROR", "No config", errors.New(""), "ME00001"

	}
	//	return true, userKeyHeader

	//-------------------------------------------------------------------------------------
	//retreive userid
	authHeader := c.Get("Authorization")
	if len(authHeader) < 10 {
		return false, me, "UNAUTHORIZED", "Missing token", errors.New(""), "ME00001"
	}
	if authHeader[:7] != "Bearer " {
		return false, me, "UNAUTHORIZED", "Invalid token format", errors.New(""), "ME00002"
	}
	token := authHeader[7:]
	matched, errReg := regexp.Match("^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$", []byte(token))
	if errReg != nil {
		return false, me, "INTERNAL_ERROR", "Cannot test token regex", errReg, "ME00003"
	}
	if !matched {
		return false, me, "INTERNAL_ERROR", "Token not matching regex", errors.New(""), "ME00004"
	}

	success, userId, errstr, errUid := GetUserIdFromToken(db, config.Server.JWT_SECRET, token)
	if !success {
		return false, me, "INTERNAL_ERROR", errstr, errUid, "ME00006"
	}

	//-------------------------------------------------------------------------------------
	//COUNT nb of revoked jwt matching
	var count int64
	{
		result := db.Model(&m.RevokedJwt{}).Where("token = ?", token).Count(&count)
		if result.Error != nil {
			return false, me, "INTERNAL_ERROR", "GetUserIdAndRightsFromHeader cannot count rights for user record", result.Error, "ME00901"
		}
	}
	//CHECK
	if count > 0 {
		return false, me, "REVOKED", "Token has been revoked", errors.New(""), "ME00600"
	}

	//FETCH person
	p := m.Person{}
	if !query.FirstWhere(r, db.Where("user_id = ?", userId), &p, "ME006-010") {
		p.UserId = &userId
		if !query.Create(r, db, &p, "ME006-010") {
			return false, me, "INTERNAL_ERROR", "Could not create person", errors.New(""), "ME006-010"
		}
		return false, me, "INTERNAL_ERROR", "No person with user_id", errors.New(""), "ME006-010"
	}
	//rights := p.Rights
	//if rights.ID == 0 {
	//	return me, "INTERNAL_ERROR", "GetUserIdAndRightsFromHeader cannot find rights for user record", errors.New(""), "ME00005"
	//}

	rights := m.Rights{}
	{
		result := db.Where("user_id = ?", userId).First(&rights)
		if result.Error != nil {
			return false, me, "INTERNAL_ERROR", "GetUserIdAndRightsFromHeader cannot find rights for user record", result.Error, "ME00005"
		}
	}

	return true, structs.Me{UserId: userId, Rights: rights.Rights, PersonId: p.ID}, "OK", "", nil, ""

}
