package business

import (

	//	"gorm.io/gorm"
	//	"stretches-common-api/config"

	"stretches-common-api/app"
	m "stretches-common-api/models"
	"stretches-common-api/mysendinblue"
)

func Resend(r app.RouteContext, emailConfirmRequest m.EmailConfirmRequest) (bool, string, error, string, bool) {
	emailConfirmRequest.Renewed = true
	dbresult2b := r.GetDb().Save(&emailConfirmRequest)
	if dbresult2b.Error != nil {
		return false, "update emailConfirmRequest record", dbresult2b.Error, "", false
	}

	//generate new request
	confirmtoken := GenerateEmailConfirmToken(emailConfirmRequest.Email)

	//send request by email
	sendSuccess := mysendinblue.TriggerRegisterAskEmailConfirmationDirect(r, emailConfirmRequest.Email, confirmtoken)

	//record request
	re := m.EmailConfirmRequest{
		Email:     emailConfirmRequest.Email,
		Token:     confirmtoken,
		Sent:      false,
		Confirmed: false,
	}
	if sendSuccess {
		re.Sent = true
	}
	dbresult4 := r.GetDb().Create(&re)
	if dbresult4.Error != nil {
		return false, "create EmailConfirmRequest record", dbresult4.Error, "", false
	}
	return true, "", nil, confirmtoken, sendSuccess
}
