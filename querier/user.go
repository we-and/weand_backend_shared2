package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchUserGeneric(r app.RouteContext, userId uint32, errCode string, f func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB)) (m.User, bool) {
	db := r.GetDb()
	user := m.User{}
	{
		result := f(db, userId, &user)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return m.User{}, false
		}
		if user.ID == 0 {
			app.SetAndSaveInternalError(r, "fetch user not found", result.Error, errCode)
			return m.User{}, false
		}
	}
	return user, true
}
func (q *Querier) FetchUser(r app.RouteContext, userId uint32, errCode string) (m.User, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB) {
		return db.Find(item, userId)
	}
	return q.FetchUserGeneric(r, userId, errCode, f)
}

func (q *Querier) FetchUserWithEmailAndPhone(r app.RouteContext, userId uint32, errCode string) (m.User, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB) {
		return db.Preload("AuthMagiclink").Preload("AuthPhone").Preload("AuthEmail").Where("").Find(item, userId)
	}
	return q.FetchUserGeneric(r, userId, errCode, f)
}

func (q *Querier) FetchUserWithEmail(r app.RouteContext, userId uint32, errCode string) (m.User, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB) {
		return db.Preload("AuthMagiclink").Preload("AuthEmail").Find(item, userId)
	}
	return q.FetchUserGeneric(r, userId, errCode, f)
}

func (q *Querier) FetchUserWithLocations(r app.RouteContext, userId uint32, errCode string) (m.User, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB) {
		return db.Preload("Locations").Find(item, userId)
	}
	return q.FetchUserGeneric(r, userId, errCode, f)
}

func (q *Querier) FetchUserWithProfile(r app.RouteContext, userId uint32, errCode string) (m.User, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.User) (tx *gorm.DB) {
		return db.Preload("Profile").Find(item, userId)
	}
	return q.FetchUserGeneric(r, userId, errCode, f)
}
