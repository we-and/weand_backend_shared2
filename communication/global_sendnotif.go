package communication

import (
	//communication "stretches-common-api/communication"
	"errors"
	"fmt"
	app "stretches-common-api/app"
	m "stretches-common-api/models"
	query "stretches-common-api/query"
	"time"

	//	"gorm.io/gorm"

	//	project "stretches-common-api/project"
	structs "stretches-common-api/structs"
)

func SendNotification(r app.RouteContext, pBatch *m.SendBatch, pPerson *m.Person, pMessage *structs.NotificationMessage) (bool, error) {
	db := r.AppContext().GetDb()
	///read pointers
	if pMessage == nil {
		return false, errors.New("no message")
	}
	message := *pMessage
	if pPerson == nil {
		return false, errors.New("no person")
	}
	person := *pPerson

	title := message.Title
	content := message.Content
	personId := person.ID
	data := map[string]string{}
	success, err := SendNotif(r, personId, title, content, db, data)
	///save job status
	job := m.SendJob{
		PersonId:    person.ID,
		TeamId:      message.TeamId,
		RelatedId:   message.RelatedId,
		Destination: "app",
		Shortmedium: "N",
		Shortobject: message.GetShortObject(),
		Success:     success,
		Errdesc:     fmt.Sprintf("%v", err),
	}
	job.Started = true
	job.Sent = true
	t := time.Now()
	job.Senddate = &t

	if pBatch != nil {
		batch := (*pBatch)
		job.BatchId = (batch).ID
		(*pBatch).Jobs = append((*pBatch).Jobs, job)
	}
	if !query.Save(r, db, &job, "001") {
		return false, errors.New("job not saved")
	}

	return success, err

}
