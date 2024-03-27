package business

import (
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

// RETURNS
// bool success
// string errStr
// error error
// bool isfound
// int32 userId
func GetContactUserId(db *gorm.DB, contactId uint32) (bool, string, error, bool, uint32) {

	//get contact
	contact := m.Contact{}
	dbresultExisting := db.Find(&contact, contactId)
	if dbresultExisting.Error != nil {
		return false, "get contact by id", dbresultExisting.Error, false, 0
	}

	contactemail := []m.ContactEmail{}
	dbresultExisting2 := db.Where("contact_id = ?", contactId).Find(&contactemail)
	if dbresultExisting2.Error != nil {
		return false, "get contactemail", dbresultExisting2.Error, false, 0
	}

	for _, v := range contactemail {
		var count int64
		dbresultExisting2 := db.Model(&m.AuthMagiclink{}).Where("email = ?", v.Email).Count(&count)
		if dbresultExisting2.Error != nil {
			return false, "get ml count", dbresultExisting2.Error, false, 0
		}

		if count > 0 {
			magiclink := m.AuthMagiclink{}
			dbresultExisting2b := db.Model(&m.AuthMagiclink{}).Where("email = ?", v.Email).First(&magiclink)
			if dbresultExisting2b.Error != nil {
				return false, "get ml", dbresultExisting2b.Error, false, 0
			}
			return true, "", nil, true, magiclink.UserId
		}
	}
	return true, "", nil, false, 0

}

func GetContactUserIdByIdentifier(db *gorm.DB, contactIdentifier string) (bool, string, error, bool, uint32, m.Contact) {

	//get contact
	contact := m.Contact{}
	dbresultExisting := db.Where("identifier = ?", contactIdentifier).First(&contact)
	if dbresultExisting.Error != nil {
		return false, "get contact by identifier", dbresultExisting.Error, false, 0, m.Contact{}
	}

	contactemail := []m.ContactEmail{}
	dbresultExisting2 := db.Where("contact_id = ?", contact.ID).Find(&contactemail)
	if dbresultExisting2.Error != nil {
		return false, "get contactemail", dbresultExisting2.Error, false, 0, contact
	}

	for _, v := range contactemail {
		var count int64
		dbresultExisting2 := db.Model(&m.AuthMagiclink{}).Where("email = ?", v.Email).Count(&count)
		if dbresultExisting2.Error != nil {
			return false, "get ml count", dbresultExisting2.Error, false, 0, contact
		}

		if count > 0 {
			magiclink := m.AuthMagiclink{}
			dbresultExisting2b := db.Model(&m.AuthMagiclink{}).Where("email = ?", v.Email).First(&magiclink)
			if dbresultExisting2b.Error != nil {
				return false, "get ml", dbresultExisting2b.Error, false, 0, contact
			}
			return true, "", nil, true, magiclink.UserId, contact
		}
	}
	return true, "", nil, false, 0, contact

}
