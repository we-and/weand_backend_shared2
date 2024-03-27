package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type NotificationSettings struct {
	Email NotificationSettingsType `json:"email"`
	SMS   NotificationSettingsType `json:"sms"`
	Notif NotificationSettingsType `json:"notif"`
}

type NotificationSettingsType struct {
	ChatMessage   bool `json:"chatmessage"`
	EventInvited  bool `json:"eventinvited"`
	MemberJoined  bool `json:"memberjoined"`
	MemberReplied bool `json:"memberreplied"`
	TeamInvited   bool `json:"teaminvited"`
	Rsvp          bool `json:"rsvp"`
}

func GetDefaultNotificationSettings() NotificationSettingsType {
	return NotificationSettingsType{
		ChatMessage:   true,
		EventInvited:  true,
		MemberJoined:  true,
		TeamInvited:   true,
		Rsvp:          true,
		MemberReplied: false,
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *NotificationSettings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := NotificationSettings{}
	err := json.Unmarshal(bytes, &result)
	*j = NotificationSettings(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j NotificationSettings) Value() (driver.Value, error) {
	return json.Marshal(j)
}
