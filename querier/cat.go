package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchCategoriesGeneric(r app.RouteContext, userIds []uint32, errCode string, f func(db *gorm.DB, userIds []uint32, item *[]m.Category) (tx *gorm.DB)) ([]m.Category, bool) {
	db := r.GetDb()
	users := []m.Category{}
	{
		result := f(db, userIds, &users)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return []m.Category{}, false
		}
	}
	return users, true
}
func (q *Querier) FetchOrgCategories(r app.RouteContext, errCode string) ([]m.Category, bool) {
	f := func(db *gorm.DB, userIds []uint32, item *[]m.Category) (tx *gorm.DB) {
		return db.Preload("LinkOrgCategories").Preload("LinkOrgCategories.Org").Preload("LinkOrgCategories.Org.LinkCategories.Category").Preload("LinkOrgCategories.Org.LinkCategories.Category").Where("type = 'ORG'").Find(item)
	}
	return q.FetchCategoriesGeneric(r, []uint32{}, errCode, f)
}
