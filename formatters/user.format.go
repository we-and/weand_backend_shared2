package formatters

import (
	"stretches-common-api/business"
	m "stretches-common-api/models"
	"stretches-common-api/publicid"
	"stretches-common-api/structs"
	"time"
)

type UserFormattedItem struct {
	FirstName     string     `json:"first_name"`
	MiddleName    string     `json:"middle_name"`
	LastName      string     `json:"last_name"`
	Identifier    string     `json:"identifier"`
	Email         string     `json:"email,omitempty"`
	TrainingTimes string     `json:"trainingtimes,omitempty"`
	Phone         string     `json:"phone,omitempty"`
	IsYou         bool       `json:"is_you,omitempty"`
	Pid           uint32     `json:"pid"`
	PersonPid     uint32     `json:"person_pid"`
	Id            uint32     `json:"id"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`

	Profile            ProfileFormattedItem          `json:"profile"`
	Details            m.ContactDetailsLightData     `json:"details"`
	Iterators          []IteratorFormattedItem       `json:"iterators"`
	UserAnalytics      m.UserAnalytics               `json:"useranalytics"`
	Device             []DeviceFormattedItem         `json:"devices,omitempty"`
	Profiles           []WorkoutprofileFormattedItem `json:"profiles,omitempty"`
	AuthStrategies     []AuthstrategyFormattedItem   `json:"authstrategies,omitempty"`
	Loginattempts      []LoginattemptsFormattedItem  `json:"loginattempts,omitempty"`
	Requests           []RequestFormattedItem        `json:"requests,omitempty"`
	GptChats           []GptChatFormattedItem        `json:"gptchats,omitempty"`
	ActivePlan         PlanFormattedItem             `json:"active_plan,omitempty"`
	ActiveSubscription SubscriptionFormattedItem     `json:"active_subscription,omitempty"`
}
type ProgressWorkoutFormattedItem struct {
	ID          uint32    `json:"id"`
	WorkoutId   uint32    `json:"workout_id"`
	InstanceId  string    `json:"instance_id,omitempty"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	DurationSec int       `json:"duration_sec,omitempty"`
	EndedAt     time.Time `json:"ended_at,omitempty"`
	EndType     string    `json:"ended_type,omitempty"`
}

func FormatProgressWorkouts(items []m.ProgressWorkout) []ProgressWorkoutFormattedItem {
	res := []ProgressWorkoutFormattedItem{}
	for _, v := range items {
		res = append(res, FormatProgressWorkout(v))
	}
	return res
}
func FormatProgressWorkout(item m.ProgressWorkout) ProgressWorkoutFormattedItem {
	res := ProgressWorkoutFormattedItem{
		ID:          item.ID,
		WorkoutId:   item.WorkoutId,
		InstanceId:  item.InstanceId,
		StartedAt:   item.StartedAt,
		DurationSec: item.DurationSec,
		EndedAt:     item.EndedAt,
		EndType:     item.EndType,
	}
	return res
}

type UserKeyFormattedItem struct {
	UserKey string `json:"user_key"`
	Pid     uint32 `json:"pid"`
	Id      uint32 `json:"id"`

	ProgressWorkouts []ProgressWorkoutFormattedItem `json:"progressworkouts"`
	IsDebug          bool                           `json:"is_debug"`
	CreatedAt        *time.Time                     `json:"created_at,omitempty"`
	TrainingTimes    string                         `json:"trainingtimes,omitempty"`
	Profiles         []WorkoutprofileFormattedItem  `json:"profiles"`
	// Details m.ContactDetailsLightData `json:"details"`
	Iterators     []IteratorFormattedItem `json:"iterators"`
	FCMTokens     []FCMTokenFormattedItem `json:"fcmtokens"`
	UserAnalytics m.UserAnalytics         `json:"useranalytics"`
	Device        []DeviceFormattedItem   `json:"devices,omitempty"`
	// AuthStrategies     []AuthstrategyFormattedItem  `json:"authstrategies,omitempty"`
	// Loginattempts      []LoginattemptsFormattedItem `json:"loginattempts,omitempty"`
	// Requests []RequestFormattedItem `json:"requests,omitempty"`
	GptChats []GptChatFormattedItem `json:"gptchats,omitempty"`
	// ActivePlan         PlanFormattedItem            `json:"active_plan,omitempty"`
	// ActiveSubscription SubscriptionFormattedItem    `json:"active_subscription,omitempty"`
}

func FormatUsers(items []m.User, me structs.Me) []UserFormattedItem {
	res := []UserFormattedItem{}
	for _, v := range items {
		res = append(res, FormatUser(v, me))
	}
	return res
}
func FormatUsersForAdmin(items []m.User, me structs.Me, langCode string) []UserFormattedItem {
	res := []UserFormattedItem{}
	for _, v := range items {
		res = append(res, FormatUserForAdmin(v, me, langCode))
	}
	return res
}
func FormatUserKeysForAdmin(items []m.UserKey, langCode string) []UserKeyFormattedItem {
	res := []UserKeyFormattedItem{}
	for _, v := range items {
		res = append(res, FormatUserKeyForAdmin(v, langCode ))
	}
	return res
}
func FormatUser(u m.User, me structs.Me) UserFormattedItem {
	res := UserFormattedItem{
		FirstName:          u.FirstName,
		MiddleName:         "",
		CreatedAt:          u.CreatedAt,
		LastName:           u.LastName,
		Email:              u.GetEmail(),
		ActivePlan:         FormatPlan(u.ActivePlan),
		ActiveSubscription: FormatSubscription(u.ActiveSubscription),
		Phone:              business.GetUserPhone(u),
		//	Details:            m.ContactDetailsLightData{Emails: []string{item.RetreiveEmail()}},
		Pid:   publicid.Obfuscate32bit(u.ID),
		IsYou: me.CheckUserId(u.ID),
	}
	res.Device = FormatDevices((u.Devices))
	//	if u.GptChats != nil {

	//	}
	if u.Profile != nil {
		res.Profile = FormatProfile(*u.Profile)
	}
	return res
}

// same with real id
func FormatUserForAdmin(u m.User, me structs.Me, langCode string) UserFormattedItem {
	res := UserFormattedItem{
		FirstName:          u.FirstName,
		MiddleName:         "",
		LastName:           u.LastName,
		CreatedAt:          u.CreatedAt,
		Loginattempts:      FormatLoginattempts(u.LoginAttempts),
		AuthStrategies:     FormatAuthstrategies(u.AuthStrategies),
		Email:              u.GetEmail(),
		ActivePlan:         FormatPlan(u.ActivePlan),
		ActiveSubscription: FormatSubscription(u.ActiveSubscription),
		Phone:              business.GetUserPhone(u),
		Pid:                publicid.Obfuscate32bit(u.ID),
		Id:                 (u.ID),
		IsYou:              me.CheckUserId(u.ID),
	}
	var trainingtime string
	err := u.TrainingTimes.Recap.AssignTo(&trainingtime)
	if err != nil {
	}
	res.Iterators = FormatIteratorsForAdmin(u.Iterators,langCode)

	res.TrainingTimes = trainingtime
	res.UserAnalytics = u.UserAnalytics
	res.Device = FormatDevices((u.Devices))
	res.Requests = FormatRequests(u)
	res.GptChats = FormatGptChats((u.GptChats))
	if u.Profile != nil {
		res.Profile = FormatProfile(*u.Profile)
	}
	return res
}

// same with real id
func FormatUserKeyForAdmin(u m.UserKey, langCode string) UserKeyFormattedItem {
	res := UserKeyFormattedItem{
		CreatedAt: u.CreatedAt,
		UserKey:   u.UserKey,
		IsDebug:   u.IsDebug,
		Pid:       publicid.Obfuscate32bit(u.ID),
		Id:        (u.ID),
	}
	var trainingtime string
	err := u.TrainingTimes.Recap.AssignTo(&trainingtime)
	if err != nil {
	}
	res.ProgressWorkouts = FormatProgressWorkouts(u.ProgressWorkouts)
	//	res.Profiles = FormatWorkoutprofiles(u.Profiles)
	res.TrainingTimes = trainingtime
	res.UserAnalytics = u.UserAnalytics
	res.Iterators = FormatIteratorsForAdmin(u.Iterators, langCode )
	res.Device = FormatLinkDeviceUsers((u.Devices))
	res.FCMTokens = FormatFCMTokens((u.FCMTokens))
	res.GptChats = FormatGptChats((u.GptChats))
	return res
}
func FormatUserWithProfile(item m.User, me structs.Me, profile m.Profile) UserFormattedItem {
	return UserFormattedItem{
		FirstName:  item.FirstName,
		MiddleName: "",
		LastName:   item.LastName,

		ActivePlan:         FormatPlan(item.ActivePlan),
		ActiveSubscription: FormatSubscription(item.ActiveSubscription),
		Email:              item.GetEmail(),
		Profile:            FormatProfile(profile),
		Phone:              business.GetUserPhone(item),
		Details:            m.ContactDetailsLightData{Emails: []string{item.RetreiveEmail()}},
		Pid:                publicid.Obfuscate32bit(item.ID),
		IsYou:              me.CheckUserId(item.ID),
	}
}

type WorkoutprofileFormattedItem struct {
	Sports     []string `json:"sports"`
	Injuries   []string `json:"injuries"`
	Focus      []string `json:"focus"`
	Conditions []string `json:"conditions"`
	Goals      []string `json:"goals"`
	Gear       []string `json:"gear"`
	Experience []string `json:"experience"`
}

func FormatWorkoutprofiles(items []m.LinkUserWorkoutprofile) []WorkoutprofileFormattedItem {
	res := []WorkoutprofileFormattedItem{}
	for _, v := range items {
		if v.WorkoutProfileRaw == nil {
			res = append(res, FormatWorkoutprofile(*v.WorkoutProfileRaw))
		}
	}
	return res
}
func FormatWorkoutprofile(item m.WorkoutProfileRaw) WorkoutprofileFormattedItem {
	res := WorkoutprofileFormattedItem{}
	return res
}
