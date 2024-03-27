package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchUsersGeneric(r app.RouteContext, userIds []uint32, errCode string, f func(db *gorm.DB, userIds []uint32, item *[]m.User) (tx *gorm.DB)) ([]m.User, bool) {
	db := r.GetDb()
	users := []m.User{}
	{
		result := f(db, userIds, &users)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return []m.User{}, false
		}
	}
	return users, true
}
func (q *Querier) FetchUsersInListWithProfileAndEmailAndPhone(r app.RouteContext, userIds []uint32, errCode string) ([]m.User, bool) {
	f := func(db *gorm.DB, userIds []uint32, item *[]m.User) (tx *gorm.DB) {
		return db.Preload("AuthPhone").Preload("AuthMagiclink").Preload("Profile").Where("id IN ?", userIds).Find(item)
	}
	return q.FetchUsersGeneric(r, userIds, errCode, f)
}
