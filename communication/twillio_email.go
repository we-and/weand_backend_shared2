package communication

import (
	"errors"
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendTwilioEmail(twillioCredentials *TwillioCredentials, templateId string, fromAddress string, fromName string, toEmail string, dynamicMap map[string]string) (resSuccess bool, resError error) {
	if twillioCredentials == nil {
		resSuccess = false
		resError = errors.New("no twilio settings")
		return
	}
	m := mail.NewV3Mail()

	//from
	e := mail.NewEmail(fromName, fromAddress)
	m.SetFrom(e)
	m.SetTemplateID(templateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail(toEmail, toEmail),
	}
	p.AddTos(tos...)

	for k, v := range dynamicMap {
		p.SetDynamicTemplateData(k, v)
	}
	m.AddPersonalizations(p)

	request := sendgrid.GetRequest(*&twillioCredentials.EmailApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(m)
	request.Body = Body
	response, err := sendgrid.API(request)

	if err != nil {
		return false, err
	} else {
		if response != nil {
			code := (*response).StatusCode
			if code == 403 {
				resSuccess = false
				resError = errors.New((*response).Body)
				return
			} else if code == 401 {
				resSuccess = false
				resError = errors.New((*response).Body)
				return
			} else {
				fmt.Printf("Response code=%v body=%v", code, (*response).Body)
				resSuccess = true
				resError = nil
				return
			}
		} else {
			resSuccess = false

			resError = errors.New("no response from twilio")
			return
		}

	}

}
