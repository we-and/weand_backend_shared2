package test

import (
	config "stretches-common-api/config"
	"strings"
)

func IsRealUser(email string) bool {
	isBypassEmail := IsBypassEmail(email)
	isInternalEmail := IsInternalTest(email)
	isAutotestEmail := IsAutoTest(email)
	isStaffEmail := IsStaff(email)
	isGreylistedEmail := IsGreylistedEmail(email)
	return !isBypassEmail && !isInternalEmail && !isAutotestEmail && !isGreylistedEmail && !isStaffEmail
}
func IsTestDevice(deviceId string) bool {
	testDeviceIds := []string{
		"04A9EA98-7F9F-4FED-822E-8825BB3783D7",
	}

	for _, v := range testDeviceIds {
		if deviceId == v {
			return true
		}
	}
	return false
}
func IsAutoTest(email string) bool {
	return strings.Contains(email, "@weand.co.uk") && strings.Contains(email, "autotest-")
} // {
func IsStaff(email string) bool {
	return strings.Contains(email, "@weand.co.uk")
} // {
func IsInternalTest(email string) bool {
	return strings.Contains(email, "@weand.co.uk") && strings.Contains(email, "internal-")
} // {
func IsBypassEmail(email string) bool {
	return email == "bypass-apple@weand.co.uk" || email == "bypass-google@weand.co.uk"
} // {
func IsGreylistedEmail(email string) bool {
	return email == "test@google.com"
} // {

func OverwriteEmailIfTest(email string, config config.AppConfig) string {
	if IsInternalTest(email) || IsAutoTest(email) {
		return config.AddressForRedictedTestEmails
	}
	return email
}
