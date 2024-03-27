package feed

import (
	m "stretches-common-api/models"

	"gorm.io/gorm"
)

func CreateFeedMemberJoinedTeam(userId uint32, teamId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(userId, "MEMBER_JOINED_TEAM", teamId, 0, db)
}
func CreateFeedCoachJoinedTeam(creatorUserId uint32, teamId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(creatorUserId, "COACH_JOINED_TEAM", teamId, 0, db)
}
func CreateFeedTeamCreated(teamId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(0, "TEAM_CREATED", teamId, 0, db)
}
func CreateFeedTeamRenamed(teamId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(0, "TEAM_RENAMED", teamId, 0, db)
}
func CreateFeedEventCreated(creatorUserId uint32, teamId uint32, relatedId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(creatorUserId, "EVENT_CREATED", teamId, relatedId, db)
}
func CreateFeedPersonAdded(creatorUserId uint32, teamId uint32, relatedId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(creatorUserId, "PERSON_ADDED", teamId, relatedId, db)
}
func CreateFeedEventRenamed(creatorUserId uint32, teamId uint32, relatedId uint32, db *gorm.DB) (bool, string, error) {
	return CreateRelatedFeed(creatorUserId, "EVENT_RENAMED", teamId, relatedId, db)
}
func CreateFeedEventReplied(creatorUserId uint32, teamId uint32, reply string, eventId uint32, occurenceTimestamp int64, db *gorm.DB) (bool, string, error, m.Activity) {
	return CreateFeedWithAnotherRelatedOccurence(creatorUserId, "EVENT_REPLIED", reply, teamId, eventId, occurenceTimestamp, db)
}
func CreateFeedPollReplied(creatorUserId uint32, teamId uint32, reply string, pollId uint32, db *gorm.DB) (bool, string, error, m.Activity) {
	return CreateFeedWithAnotherRelated(creatorUserId, "POLL_REPLIED", reply, teamId, pollId, db)
}
func CreateFeed(userId uint32, type_ string, subtype string, relatedId uint32, db *gorm.DB) (bool, string, error) {
	feed := m.Activity{
		UserId: userId,
		Type_:  type_,
	}
	result3 := db.Create(&feed)
	if result3.Error != nil {
		return false, "save activity", result3.Error
	}
	return true, "", nil
}
func CreateRelatedFeed(userId uint32, type_ string, teamId uint32, relatedId uint32, db *gorm.DB) (bool, string, error) {
	feed := m.Activity{
		Type_:  type_,
		TeamId: teamId,
	}
	if userId > 0 {
		feed.UserId = userId
	}
	if relatedId > 0 {
		feed.RelatedId = relatedId
	}
	result3 := db.Create(&feed)
	if result3.Error != nil {
		return false, "save activity", result3.Error
	}
	return true, "", nil
}

func CreateFeedWithAnotherRelatedOccurence(userId uint32, type_ string, subtype string, teamId uint32, relatedId uint32, related2Id int64, db *gorm.DB) (bool, string, error, m.Activity) {
	feed := m.Activity{
		UserId:      userId,
		Type_:       type_,
		Subtype:     subtype,
		TeamId:      teamId,
		RelatedId:   relatedId,
		OccurenceId: related2Id,
	}
	result3 := db.Create(&feed)
	if result3.Error != nil {
		return false, "save activity", result3.Error, m.Activity{}
	}
	return true, "", nil, feed
}
func CreateFeedWithAnotherRelated(userId uint32, type_ string, subtype string, teamId uint32, relatedId uint32, db *gorm.DB) (bool, string, error, m.Activity) {
	feed := m.Activity{
		UserId:    userId,
		Type_:     type_,
		Subtype:   subtype,
		TeamId:    teamId,
		RelatedId: relatedId,
	}
	result3 := db.Create(&feed)
	if result3.Error != nil {
		return false, "save activity", result3.Error, m.Activity{}
	}
	return true, "", nil, feed
}
func CreateFeedWithExtra(userId uint32, type_ string, subtype string, relatedId uint32, extra string, db *gorm.DB) (bool, string, error) {
	feed := m.Activity{
		UserId: userId,
		/*		RelatedId: relatedId,
				Extra:extra,
				Type_:     type_,
				Subtype:   subtype,
		*/}
	result3 := db.Create(&feed)
	if result3.Error != nil {
		return false, "save activity", result3.Error
	}
	return true, "", nil
}
