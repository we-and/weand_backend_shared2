package business

import (
	"stretches-common-api/models"
)

func ExtractContactIdsFromContactEmails(contactemails []models.ContactEmail) []uint32 {
	contactMap := map[uint32]bool{}
	contactIds := []uint32{}
	for i := 0; i < len(contactemails); i++ {
		ce := contactemails[i]
		contactMap[ce.ContactId] = true
	}
	for key, _ := range contactMap {
		contactIds = append(contactIds, key)
	}
	return contactIds
}

func ExtractContactIdsFromContact(contacts []models.Contact) []uint32 {
	contactMap := map[uint32]bool{}
	contactIds := []uint32{}
	for i := 0; i < len(contacts); i++ {
		ce := contacts[i]
		contactMap[ce.ID] = true
	}
	for key, _ := range contactMap {
		contactIds = append(contactIds, key)
	}
	return contactIds
}
