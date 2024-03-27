package business

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"
	"stretches-common-api/query"

	"gorm.io/gorm"
)

func GenerateNewFreeSubscription(r app.RouteContext, db *gorm.DB, user m.User) (resSuccess bool) {
	//--------------------------------------------------------------------------------------
	//CREATE Subscription
	sub := m.Subscription{PricingId: 5, UserId: user.ID}
	if !query.CreateOrRollback(r, db, &sub, "ME004-303") {
		return
	}

	//--------------------------------------------------------------------------------------
	//FETCH plan
	plan := m.Plan{}
	if !query.FirstWhereOrRollback(r, db.Where("name = ?", "FREE"), &plan, "ME004-304") {
		return
	}

	changes := map[string]interface{}{

		"active_subscription_id": sub.ID,
		"active_plan_id":         plan.ID,
	}
	if !query.UpdateWhere(r, db.Model(&m.User{}).Where("id = ?", user.ID), &changes, "ME0890") {
		return
	}
	resSuccess = true
	return
}
