package business

import (
	"stretches-common-api/models"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

type MatchingUserItem struct {
	Type    string      `json:"type"`
	Matched string      `json:"matched"`
	UserId  uint32      `json:"user_id"`
	User    models.User `json:"user"`
}

// create a monyl account
func FindMonylUser(db *gorm.DB, details m.ContactDetailsLightData) []MatchingUserItem {
	res := []MatchingUserItem{}
	for _, email := range details.Emails {
		var count int64
		{
			result := db.Model(&m.AuthMagiclink{}).
				Where("email = ? ", email).Count(&count)
			if result.Error != nil {
				//return false, "find fcmtoken", result.Error
			}
		}
		if count > 0 {
			//FETCH magiclink
			item := m.AuthMagiclink{}
			{
				result := db.Where("email = ? ", email).First(&item)
				if result.Error != nil {
					//return false, "find fcmtoken", result.Error
				}
			}
			//FETCH magiclink
			user := m.User{}
			{
				result := db.Preload("AuthPhone").Find(&user, item.UserId)
				if result.Error != nil {
					//return false, "find fcmtoken", result.Error
				}
			}
			user.AuthMagiclink = item
			res = append(res, MatchingUserItem{
				Type:    "EMAIL",
				Matched: email,
				UserId:  item.UserId,
				User:    user,
			})
		}
	}
	for _, fullphone := range details.Phones {

		phone, countryCode := ExtractPhoneCountryCode(fullphone)
		var count int64
		{
			result := db.Model(&m.AuthPhone{}).
				Where("phone = ? AND country_code = ? AND confirmed = true", phone, countryCode).Count(&count)
			if result.Error != nil {
				//return false, "find fcmtoken", result.Error
			}
		}
		if count > 0 {
			item := m.AuthPhone{}
			{
				result := db.
					Where("phone = ? AND country_code = ? AND confirmed = true", phone, countryCode).First(&item)
				if result.Error != nil {
					//return false, "find fcmtoken", result.Error
				}
			}
			//FETCH magiclink
			user := m.User{}
			{
				result := db.Preload("AuthMagiclink").Find(&user, item.UserId)
				if result.Error != nil {
					//return false, "find fcmtoken", result.Error
				}
			}
			user.AuthPhone = item

			res = append(res, MatchingUserItem{
				Type:    "PHONE",
				Matched: fullphone,
				UserId:  item.UserId,
				User:    user,
			})
		}
		isAlreadyIn := false
		userIds := map[uint32]bool{}
		for _, matching := range res {

			for id, _ := range userIds {
				if matching.UserId == id {
					isAlreadyIn = true
					break
				}
			}
			if !isAlreadyIn {
				res = append(res, MatchingUserItem{
					Type:    "PHONE",
					Matched: matching.Matched,
					UserId:  matching.UserId,
					User:    matching.User,
				})
			}
			userIds[matching.UserId] = true
		}

	}
	return res

}

func ExtractPhoneCountryCode(fullphone string) (string, string) {
	if len(fullphone) == 12 {
		fpbytes := []byte(fullphone)
		ind := fpbytes[0:2]
		if string(ind) == "+1" {
			base := fpbytes[2:]
			return string(base), string(ind)
		}
	}
	return fullphone, ""

}
