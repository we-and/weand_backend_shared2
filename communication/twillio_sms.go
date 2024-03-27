package communication

import (
	"errors"
	"fmt"

	//"net/http"
	//	"net/url"
	//	"strings"
	"stretches-common-api/app"

	twilio_ "github.com/kevinburke/twilio-go"
)

type TwillioCredentials struct {
	SmsSid      string
	SmsToken    string
	SmsFrom     string
	EmailApiKey string
}

func BuildTwilio(sid string, token string, from string, emailApiKey string) TwillioCredentials {
	return TwillioCredentials{SmsSid: sid, SmsToken: token, SmsFrom: from, EmailApiKey: emailApiKey}
}

func SendSMSWithClient(r app.RouteContext, twillioCredentials *TwillioCredentials, to string, content string) (bool, error) {
	if twillioCredentials == nil {
		return false, errors.New("No twilio settings")
	}
	sid := (*twillioCredentials).SmsSid
	token := (*twillioCredentials).SmsToken
	from := (*twillioCredentials).SmsFrom

	client := twilio_.NewClient(sid, token, nil)
	_, err := client.Messages.SendMessage(from, to, content, nil)
	app.SaveInternalError(r, fmt.Sprintf("Error sending sms to %v", to), err, "ME01701")
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

