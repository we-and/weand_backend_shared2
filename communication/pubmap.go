package communication

import (
	structs "stretches-common-api/structs"
	// "fmt"
)
func GeneratePubMap(baseMessage *structs.Message, smsContent string, emailTitle string, emailContent string, emailHeaderUrl string, emailButtons []structs.EmailButton, notifTitle string, notifContent string) map[string]interface{} {
	pubMap := map[string]interface{}{}

	//SMS
	smsMsg := structs.SMSMessage{}
	smsMsg.CopyFrom(baseMessage)
	smsMsg.Content = smsContent
	pubMap["SMS"] = smsMsg

	//email
	emailMsg := structs.EmailMessage{}
	emailMsg.CopyFrom(baseMessage)
	emailMsg.Content = emailContent
	emailMsg.Title = emailTitle
	emailMsg.Buttons = emailButtons
	emailMsg.HeaderUrl = emailHeaderUrl
	pubMap["EMAIL"] = emailMsg

	//notif
	notifMsg := structs.NotificationMessage{}
	notifMsg.CopyFrom(baseMessage)
	notifMsg.Content = notifContent
	notifMsg.Title = notifTitle
	pubMap["NOTIF"] = notifMsg
	return pubMap
}
