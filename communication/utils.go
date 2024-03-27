package communication

import (
	structs "stretches-common-api/structs"
)

func GetEmailTemplateId(message *structs.EmailMessage) string {
	msg := (*message)
	if msg.IsAnnouncement() {
		return "d-8e5bd1bc303c42478d93ec97e62db2e1"
	}
	nbButtons := len(msg.Buttons)
	switch nbButtons {
	case 0:
		return "d-8e5bd1bc303c42478d93ec97e62db2e1"
	case 1:
		return "d-7a9111b772d24e6c9fa7d289b2c28e7f"
	case 2:
		return "d-cfa989b051b44dfdab6fd53ead1c0a7a"
	case 3:
		return "d-c8432227839142a1b2704c82cae5e488"
	case 4:
		return "d-f86e41eeb65e4553bd9f976d6adfa82b"
	}
	return "d-deb583dda50b4de984863b2e344a2686"

}
