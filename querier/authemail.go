package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchAuthEmailGeneric(r app.RouteContext, userId uint32, errCode string, f func(db *gorm.DB, userId uint32, item *m.AuthEmailpassword) (tx *gorm.DB)) (m.AuthEmailpassword, bool) {
	//h := r.AppCtx
	db := r.GetDb()
	obj := m.AuthEmailpassword{}
	{
		result := f(db, userId, &obj)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return m.AuthEmailpassword{}, false
		}
		if obj.ID == 0 {
			app.SetAndSaveInternalError(r, "fetch user not found", result.Error, errCode)
			return m.AuthEmailpassword{}, false
		}
	}
	return obj, true
}
func (q *Querier) FetchAuthEmail(r app.RouteContext, email string, errCode string) (m.AuthEmailpassword, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.AuthEmailpassword) (tx *gorm.DB) {
		return db.Where("email = ?  ", email).Find(item)
	}
	return q.FetchAuthEmailGeneric(r, 0, errCode, f)
}
