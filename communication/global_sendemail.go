package communication

import (
	//communication "stretches-common-api/communication"
	"errors"
	"fmt"
	app "stretches-common-api/app"
	m "stretches-common-api/models"
	query "stretches-common-api/query"
	"stretches-common-api/test"
	"time"

	//	"gorm.io/gorm"

	//	project "stretches-common-api/project"
	structs "stretches-common-api/structs"
)

func SendEmail(r app.RouteContext, pBatch *m.SendBatch, pPerson *m.Person, to string, pMessage *structs.EmailMessage) (bool, error) {
	///read pointers
	if pMessage == nil {
		return false, errors.New("no message")
	}
	message := *pMessage
	if pPerson == nil {
		return false, errors.New("no person")
	}
	person := *pPerson

	appConfig := r.AppCtx.GetConfig()
	credentials := BuildTwilioCredentials(appConfig)
	templateId := GetEmailTemplateId(pMessage)
	destinationEmail := test.OverwriteEmailIfTest(to, r.GetConfigA())

	fromAddress := "helloakto@weand.co.uk"
	fromName := message.TeamName //"TeamReturn"
	dynamicMap := map[string]string{
		"title":   message.Title,
		"content": message.Content,
	}
	nbButtons := len(message.Buttons)
	if nbButtons > 0 {
		b := message.Buttons[0]
		dynamicMap["link"] = b.Link
		dynamicMap["label"] = b.Label
	}
	if nbButtons > 1 {
		b := message.Buttons[1]
		dynamicMap["link2"] = b.Link
		dynamicMap["label2"] = b.Label
	}
	if nbButtons > 2 {
		b := message.Buttons[2]
		dynamicMap["link3"] = b.Link
		dynamicMap["label3"] = b.Label
	}
	if nbButtons > 3 {
		b := message.Buttons[3]
		dynamicMap["link4"] = b.Link
		dynamicMap["label4"] = b.Label
	}
	success, err := SendTwilioEmail(&credentials, templateId, fromAddress, fromName, destinationEmail, dynamicMap)
	job := m.SendJob{
		PersonId:    person.ID,
		TeamId:      message.TeamId,
		RelatedId:   message.RelatedId,
		Destination: to,
		Shortmedium: "E",
		Shortobject: message.GetShortObject(),
		Success:     success,
		Errdesc:     fmt.Sprintf("%v", err),
	}
	job.Started = true
	job.Sent = true
	t := time.Now()
	job.Senddate = &t

	if pBatch != nil {
		job.BatchId = (*pBatch).ID
		(*pBatch).Jobs = append((*pBatch).Jobs, job)

	}
	db := r.AppContext().GetDb()
	if !query.Save(r, db, &job, "001") {
		return false, errors.New("job not saved")
	}

	return success, err
}
