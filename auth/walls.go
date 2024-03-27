package auth

import (
	"errors"
	"stretches-common-api/app"
	"stretches-common-api/structs"
)

// / Reads userId from request header "Authorization" and checks if user has ADMIN rights
// / returns: int userId if user exists and has ADMIN rights
// /          -1 otherwise
func AdminWall(r app.RouteContext) (bool, structs.Me) {
	//read token for userid and rights
	success, accessDetails, errtype, errstr, err, errCode := GetUserIdAndRightsFromHeader(r)
	if !success {
		app.SetError(r, errstr, err, errtype, errCode)
		return false, accessDetails
	}

	//check rights
	if !HasAdminRights(accessDetails.Rights) {
		app.SetUnauthorized(r, "Access forbidden. User only. ", errors.New(""), errCode)
		return false, accessDetails
	}
	return true, accessDetails
}

// / Reads userId from request header "Authorization" and checks if user has USER rights
// / returns: int userId if user exists and has USER rights
// /          -1 otherwise
func UserWall(r app.RouteContext) (bool, structs.Me) {
	//read token for userid and rights
	success, accessDetails, errtype, errstr, err, errCode := GetUserIdAndRightsFromHeader(r)
	if !success {
		app.SetError(r, errstr, err, errtype, errCode)
		return false, accessDetails
	}

	//check rights
	if !HasUserRights(accessDetails.Rights) {
		app.SetUnauthorized(r, "Access forbidden. Admin only.", errors.New(""), errCode)
		return false, accessDetails
	}

	return true, accessDetails
}
