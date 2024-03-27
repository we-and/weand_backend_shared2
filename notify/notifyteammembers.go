package notify

import (
	"stretches-common-api/app"
	m "stretches-common-api/models"

	"stretches-common-api/communication"
	"stretches-common-api/query"

	"gorm.io/gorm"
)

func NotifFilterNotifMemberJoined(u m.Person) bool {
	return u.Notificationsettings.Notif.MemberJoined
}
func NotifFilterNotifMemberReplied(u m.Person) bool {
	return u.Notificationsettings.Notif.MemberReplied
}
func NotifyTeamMembers(r app.RouteContext, db *gorm.DB, meUserId uint32, team m.Team, title string, msg string, userFilter func(m.Person) bool, data map[string]string) bool {
	//FETCH team members
	seats := []m.Seat{}
	if !query.FindWhere(r, db.Preload("Person").Preload("Person.User").Preload("Person.Unregistered").Where("team_id = ?", team.ID), &seats, "MC001-009") {
		return false
	}
	//for each, send notif
	for _, s := range seats {
		if s.Person != nil {
			p := s.Person
			if p.IsMe(meUserId) && p.User != nil { //do not notify yourself
				filterResult := userFilter(*(p))
				if filterResult {
					communication.SendNotif(r, (p.ID), title, msg, db, data)
				}
			}
		}
	}
	return true
}
