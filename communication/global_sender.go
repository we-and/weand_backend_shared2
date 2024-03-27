package communication

import (
	//communication "stretches-common-api/communication"
	"errors"
	"fmt"
	app "stretches-common-api/app"
	m "stretches-common-api/models"
	"stretches-common-api/usage"

	//	"gorm.io/gorm"

	//	project "stretches-common-api/project"
	structs "stretches-common-api/structs"
)

/*
	func SendToPerson(r app.RouteContext, pBatch *m.SendBatch, type_ string, pTeam *m.Team, pPerson *m.Person, title string, content string, link string) (bool, error, []structs.SendRecapPerson) {
																				recap := []structs.SendRecapPerson{}
																				if pTeam == nil || pPerson == nil {
																					return false, errors.New("no team or person"), recap
		}
																				team := *pTeam
																				person := *pPerson

		message := structs.Message{
																					ActionLink: link,
																					Title:      title,
																					Content:    content,
																					TeamName:   team.Name,
																					Object:       type_,
																					TeamId:     team.ID,
		}

		details := person.Details
			details := person.Details
			emails := details.Emails
			//if false {
			for _, e := range emails {
				email := e.Email
				success, err := SendEmail(r, pBatch, pPerson, email, &message)
																				}
																				//}/
																				phones := details.Phones
																				for _, e := range phones {
																					destPhonenumber := e.Build()
																					content = fmt.Sprintf("%v %v", content, link)
																					success, err := SendSMS(r, pBatch, pPerson, destPhonenumber, &message)
																					recap = append(recap, structs.SendRecapPerson{Type: "SMS", Destination: destPhonenumber, Success: success, Error: err})

		}
																				if person.IsRegistered() {
																					success, err := SendNotification(r, pBatch, pPerson, &message)
																					recap = append(recap, structs.SendRecapPerson{Type: "NOTIF", Destination: fmt.Sprintf("%v", person.ID), Success: success, Error: err})

		}
																				return true, nil, recap
																			}
*/
func SendToPerson2(r app.RouteContext, teamId uint32, pBatch *m.SendBatch, pPerson *m.Person, pubMap map[string]interface{}) (bool, error, []structs.SendRecapPerson) {
	recap := []structs.SendRecapPerson{}
	if pPerson == nil {
		return false, errors.New("no team or person"), recap
	}
	person := *pPerson
	details := person.Details
	emails := details.Emails
	nbSMS := 0
	nbEmails := 0
	nbNotifs := 0
	if msg, okMap := pubMap["EMAIL"]; okMap {
		if emailMessage, ok := msg.(structs.EmailMessage); ok {
			for _, e := range emails {
				destEmail := e.Email
				success, err := SendEmail(r, pBatch, pPerson, destEmail, &emailMessage)
				recap = append(recap, structs.SendRecapPerson{Type: "EMAIL", Destination: destEmail, Success: success, Error: err})
				if success {
					nbEmails = nbEmails + 1
				}
			}
		}
	}
	phones := details.Phones
	if msg, okMap := pubMap["SMS"]; okMap {
		if smsMessage, ok := msg.(structs.SMSMessage); ok {
			for _, e := range phones {
				phone := e.Build()
				success, err := SendSMS(r, pBatch, pPerson, phone, &smsMessage)
				recap = append(recap, structs.SendRecapPerson{Type: "SMS", Destination: phone, Success: success, Error: err})
				if success {
					nbSMS = nbSMS + 1
				}
			}
		}
	}
	if person.IsRegistered() {
		if msg, okMap := pubMap["NOTIF"]; okMap {
			if notifMessage, ok := msg.(structs.NotificationMessage); ok {
				success, err := SendNotification(r, pBatch, pPerson, &notifMessage)
				recap = append(recap, structs.SendRecapPerson{Type: "NOTIF", Destination: fmt.Sprintf("%v", person.ID), Success: success, Error: err})
				if success {
					nbNotifs = nbNotifs + 1
				}

			}
		}
	}
	usage.IncrementSendJobDetailed(r, teamId, nbSMS, nbEmails, nbNotifs)
	return true, nil, recap
}
