package communication

import (
	"context"
	"fmt"
	"stretches-common-api/app"
	m "stretches-common-api/models"

	firebasemessaging "firebase.google.com/go/v4/messaging"
	"gorm.io/gorm"
)

func SendNotifToTokenWithData(r app.RouteContext, token string, title string, body string, db *gorm.DB, data map[string]string) (bool, error) {
	// Create the message to be sent.
	notif := firebasemessaging.Notification{
		Body:  body,
		Title: title,
	}
	msg := &firebasemessaging.Message{
		Token:        token,
		Notification: &notif,
		Data:         data,
	}

	firebaseapp := r.Firebase()
	ctx := context.Background()

	client, err := firebaseapp.Messaging(ctx)
	if err != nil {
		return false, err
	}

	// Send the message and receive the response without retries.
	response, err2 := client.Send(ctx, msg)
	if err2 != nil {
		return false, err2
	}

	fmt.Printf("Notify %v : success = %v \n", token, response)
	return true, nil
}
func SendNotif(r app.RouteContext, personId uint32, text string, body string, db *gorm.DB, data map[string]string) (bool, error) {
	items := []m.FirebaseFCMToken{}
	{
		result := db.
			Where("person_id = ? AND is_active = true", personId).Find(&items)
		if result.Error != nil {
			return false, result.Error
		}
	}
	fmt.Printf("AppNotification nb fcmtokens= %d\n", len(items))
	issuccess := true
	for _, it := range items {
		suc, err := SendNotifToTokenWithData(r, it.Token, text, body, db, data)
		if !suc {
			fmt.Printf("%v ERR=%v\n", it.Token, err)
		}
		issuccess = issuccess && suc
	}
	return issuccess, nil
}
