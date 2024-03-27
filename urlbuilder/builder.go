package urlbuilder

import (
	"fmt"
)

func GetHost() string {
	return "https://stretches.weand.co.uk"
}

func GetAppHost() string {
	return "stretches://"
}

func GetWebPollReply(locale string, network string, pollId uint32, personId uint32, reply string) string {
	host := GetHost()
	return fmt.Sprintf("%s/%v/%v/poll/%v/reply/%v/%v", host, locale, network, pollId, personId, reply)

}
func GetWebEventReply(locale string, network string, eventId uint32, startdate string, personId uint32, reply string) string {
	host := GetHost()
	return fmt.Sprintf("%s/%v/%v/event/%v/%v/reply/%v/%v", host, locale, network, eventId, startdate, personId, reply)

}

func GetWebEmailConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetHost()
	return fmt.Sprintf("%s/user/live/v1/auth/emailpassword/confirm/%v/%v", host, encodedEmail, confirmToken)
}
func GetAppEmailConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetAppHost()
	return fmt.Sprintf("%s/emailconfirm/%v/%v", host, encodedEmail, confirmToken)
}
func GetWebDeviceConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetHost()
	return fmt.Sprintf("%s/user/live/v1/auth/device/confirm/%v/%v", host, encodedEmail, confirmToken)
}

func GetAppDeviceConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetAppHost()
	return fmt.Sprintf("%s/deviceconfirm/%v/%v", host, encodedEmail, confirmToken)
}

func GetAppLinksocialConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetAppHost()
	return fmt.Sprintf("%s/linksocialconfirm/%v/%v", host, encodedEmail, confirmToken)
}
func GetWebSocialConfirm(locale string, network string, encodedEmail, confirmToken string, returning bool) string {
	host := GetHost()
	return fmt.Sprintf("%s/user/live/v1/auth/linksocial/confirm/%v/%v", host, encodedEmail, confirmToken)
}
func GetWebMagiclinkConfirm(locale string, network string, encodedEmail, confirmToken string) string {
	host := GetHost()
	return fmt.Sprintf("%s/%v/%v/magiclink/confirm/%v", host, locale, network, encodedEmail, confirmToken)
}
