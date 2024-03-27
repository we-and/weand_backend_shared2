package structs

type SMSMessage struct {
	Message
}

func (j *SMSMessage) GetMedium() string {
	return "SMS"
}

type EmailButton struct {
	Label string
	Link  string
}

type EmailMessage struct {
	Message
	Title     string
	HeaderUrl string
	Buttons   []EmailButton
}

func (j *EmailMessage) GetMedium() string {
	return "EMAIL"
}

type NotificationMessage struct {
	Message
	Title string
}

func (j *NotificationMessage) GetMedium() string {
	return "NOTIF"
}

type WhatsappMessage struct {
	Message
}

func (j *WhatsappMessage) GetMedium() string {
	return "WHATSAPP"
}
