package querier

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func (q *Querier) FetchCategoryGeneric(r app.RouteContext, userIds []uint32, errCode string, f func(db *gorm.DB, userIds []uint32, item *m.Category) (tx *gorm.DB)) (m.Category, bool) {
	db := r.GetDb()
	item := m.Category{}
	{
		result := f(db, userIds, &item)
		if result.Error != nil {
			app.SetAndSaveInternalError(r, "fetch user", result.Error, errCode)
			return m.Category{}, false
		}
	}
	return item, true
}
func (q *Querier) FetchCategoryByName(r app.RouteContext, errCode string) (m.Category, bool) {
	f := func(db *gorm.DB, userIds []uint32, item *m.Category) (tx *gorm.DB) {
		return db.Preload("LinkOrgCategories").Preload("LinkOrgCategories.Org").Where("type = 'ORG'").Find(item)
	}
	return q.FetchCategoryGeneric(r, []uint32{}, errCode, f)
}

func (q *Querier) FetchCategoryWithOrgs(r app.RouteContext, categoryName string, errCode string) (m.Category, bool) {
	f := func(db *gorm.DB, userIds []uint32, item *m.Category) (tx *gorm.DB) {
		return db.Preload("LinkOrgCategories").Preload("LinkOrgCategories.Org").Where("name = ?", categoryName).Find(item)
	}
	return q.FetchCategoryGeneric(r, []uint32{}, errCode, f)
}
