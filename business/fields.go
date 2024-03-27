package business

import (
	"fmt"
	m "stretches-common-api/models"
)

func GetUserPhone(u m.User) string {
	if u.AuthPhone.ID > 0 {
		return fmt.Sprintf("%s%s", u.AuthPhone.CountryCode, u.AuthPhone.Phone)
	}
	return ""
}
func GetPersonPhone(p m.Person) string {
	phones:=p.Details.Phones
	if len(phones)>0{
		return phones[0].Build()			
	}
/*
	if p.IsRegistered() {
		if p.User.AuthPhone.ID > 0 {
			return fmt.Sprintf("%s%s", p.User.AuthPhone.CountryCode, p.User.AuthPhone.Phone)
		}
	} else {
		phones:=p.Unregistered.Details.Phones
		if len(phones)>0{
			return phones[0].Build()			
		}
	}
*/
	return ""
}

func GetUserPhoneBase(u m.User) string {
	if u.AuthPhone.ID > 0 {
		return u.AuthPhone.Phone
	}
	return ""
}
