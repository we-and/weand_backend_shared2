package querier

/*
func (q *Querier) FetchOrgGeneric(c *fiber.Ctx, userId uint32, errCode string, f func(db *gorm.DB, userId uint32, item *m.Org) (tx *gorm.DB)) (m.Org, bool) {
	h := (q.getHandler()).(app.AppHandler)
	db := h.GetDb()
	obj := m.Org{}
	{
		result := f(db, userId, &obj)
		if result.Error != nil {
			app.SetInternalError(c, h, "fetch org", result.Error, errCode)
			return m.Org{}, false
		}
		if obj.ID == 0 {
			app.SetInternalError(c, h, "fetch org not found", result.Error, errCode)
			return m.Org{}, false
		}
	}
	return obj, true
}
func (q *Querier) FetchOrg(c *fiber.Ctx, id uint32, errCode string) (m.Org, bool) {
	f := func(db *gorm.DB, userId uint32, item *m.Org) (tx *gorm.DB) {
		return db.Preload("LinkCategories").Preload("LinkCategories.Category").Find(item, id)
	}
	return q.FetchOrgGeneric(c, 0, errCode, f)
}
*/
