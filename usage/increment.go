package usage

import (
	app "stretches-common-api/app"
	m "stretches-common-api/models"
	"stretches-common-api/query"
	"time"
	// "stretches-common-api/timezone"
)

func IncrementSMSReceived(r app.RouteContext, teamId uint32) error {
	return IncrementUsage(r, teamId, "SR",1)
}
func IncrementAutoReply(r app.RouteContext, teamId uint32) error {
	return IncrementUsage(r, teamId, "RE",1)
}

func IncrementSendJob(r app.RouteContext, teamId uint32) error {
	return IncrementUsage(r, teamId, "JO",1)
}
func IncrementSendJobDetailed(r app.RouteContext, teamId uint32,nbSms int,nbEmails int,nbNotifs int) error {
	IncrementUsage(r, teamId, "SM",nbSms)
	IncrementUsage(r, teamId, "EM",nbEmails)
	IncrementUsage(r, teamId, "NO",nbNotifs)
return nil
}
func IncrementSendBatch(r app.RouteContext, teamId uint32) error {
	return IncrementUsage(r, teamId, "BA",1)
}

func IncrementUsage(r app.RouteContext, teamId uint32, usage string, n int) error {
	db := (r.AppCtx).GetDb()
	var countUsage int64
	now := time.Now()
//	date := now.Format(timezone.HumanDateFormat)
	datenohour:=time.Date(now.Year(),now.Month(),now.Day(),0,0,0,0,now.Location())

	if !query.CountWhere(r, db.Model(&m.UsageStat{}).Where("team_id = ? and date= ? and usage= ? ", teamId, datenohour, usage), &countUsage, "043") {
		return nil
	}
	if countUsage == 0 {
		usage := m.UsageStat{
			TeamId: teamId,
			Date:   datenohour,
			Count:  n,
			Usage:  "SR", //S(ms) R(eceived)
		}
		if !query.Create(r, db, &usage, "022") {
			return nil
		}
	} else {

		usage := m.UsageStat{}
		if !query.FirstWhere(r, db.Where("team_id = ? and date= ? and usage = ? ", teamId, datenohour, usage), &usage, "044") {
			return nil
		}
		usage.Count = usage.Count + n
		if !query.Save(r, db, &usage, "022") {
			return nil
		}
	}
		return nil
	}
