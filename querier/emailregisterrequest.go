package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchEmailRegisterRequestGeneric(r app.RouteContext, objId uint32, errCode string, f func(db *gorm.DB, userId uint32, item *m.EmailConfirmRequest) (tx *gorm.DB)) (m.EmailConfirmRequest, bool) {
	db := r.GetDb()
	obj := m.EmailConfirmRequest{}
	{
		result := f(db, objId, &obj)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return m.EmailConfirmRequest{}, false
		}
		if obj.ID == 0 {
			app.SetAndSaveInternalError(r, "fetch user not found", result.Error, errCode)
			return m.EmailConfirmRequest{}, false
		}
	}
	return obj, true
}
func (q *Querier) FetchEmailRegisterRequest(r app.RouteContext, email string, token string, errCode string) (m.EmailConfirmRequest, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.EmailConfirmRequest) (tx *gorm.DB) {
		return db.Where("email = ? AND token = ? ", email, token).Find(item)
	}
	return q.FetchEmailRegisterRequestGeneric(r, 0, errCode, f)
}
