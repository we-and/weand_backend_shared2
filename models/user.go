package models

import (
	"fmt"
	"time"
	_ "time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint32 `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`

	//DisplayName          string     `json:"display_name"`
	/////////	AuthStrategy         string     `json:"auth_strategy"`
	Locale               string `json:"locale"`
	ActiveSubscriptionId uint32 `json:"active_subscription_id"`
	ActivePlanId         uint32 `json:"active_plan_id"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//	Deleted              bool       `json:"deleted,omitempty"`

	OnboardingStage string    `json:"onboarding_stage"`
	IsConfirmed     bool      `json:"is_confirmed,omitempty"`
	Interests       Interests `json:"interests,omitempty"`

	//POPULATED FIELDS
	FirebaseToken  FirebaseFCMToken   `gorm:"foreignKey:user_id" json:"firebasetoken,omitempty"`
	PublicId       *PublicId          `gorm:"foreignKey:user_id" json:"publicid,omitempty"`
	Locations      []Location         `gorm:"foreignKey:user_id" json:"locations,omitempty"`
	AuthEmail      *AuthEmailpassword `gorm:"foreignKey:user_id" json:"credentials,omitempty"`
	AuthSocial     *AuthSocial        `gorm:"foreignKey:user_id" json:"auth_social,omitempty"`
	AuthStrategies []AuthStrategy     `gorm:"foreignKey:user_id" json:"auth_strategies,omitempty"`

	LoginAttempts []Loginhistory `gorm:"foreignKey:user_id" json:"login_history,omitempty"`
	TrainingTimes Trainingtimes  `gorm:"foreignKey:user_id" json:"trainingtimes,omitempty"`

	Iterators []Iterator `gorm:"foreignKey:user_id" json:"iterators,omitempty"`

	//EmailConfirmRequests  []EmailConfirmRequest      `gorm:"foreignKey:user_id" json:"emailconfirm,omitempty"`
	//DeviceConfirmRequests []DeviceConfirmRequest     `gorm:"foreignKey:user_id" json:"deviceconfirm,omitempty"`
	//SocialConfirmRequests []LinkSocialConfirmRequest `gorm:"foreignKey:user_id" json:"linksocialconfirm,omitempty"`

	UserAnalytics      UserAnalytics  `gorm:"foreignKey:user_id" json:"useranalytics,omitempty"`
	AuthPhone          AuthPhone      `gorm:"foreignKey:user_id" json:"phone,omitempty"`
	AuthMagiclink      AuthMagiclink  `gorm:"foreignKey:user_id" json:"magiclinkcredentials,omitempty"`
	Profile            *Profile       `gorm:"foreignKey:user_id" json:"profile,omitempty"`
	Devices            []Device       `gorm:"foreignKey:user_id" json:"device,omitempty"`
	Subscriptions      []Subscription `gorm:"foreignKey:user_id" json:"subscriptions,omitempty"`
	GptChats           []GPTChat      `gorm:"foreignKey:user_id" json:"gptchats,omitempty"`
	ActivePlan         Plan           `gorm:"foreignKey:active_plan_id" json:"active_plan,omitempty"`
	ActiveSubscription Subscription   `gorm:"foreignKey:active_subscription_id" json:"active_subscription,omitempty"`
}

func (c *User) GetStageIdx(stage string) uint32 {
	switch stage {
	case "ACCOUNT_SETUP":
		return 1
	case "TEAM_SETUP":
		return 2
	case "USER_SETUP":
		return 3
	default:
		return 0
	}
}

func (c *User) IsOnboardingStageReached(stage string) bool {
	thisStageIdx := c.GetStageIdx(c.OnboardingStage)
	askedStageIdx := c.GetStageIdx(stage)
	return askedStageIdx <= thisStageIdx
}

func (c *User) GetId() uint32 {
	return c.ID
}
func (c *User) GetName() string {

	if len(c.MiddleName) > 0 {
		return fmt.Sprintf("%v %v %v", c.FirstName, c.MiddleName, c.LastName)
	} else {
		return fmt.Sprintf("%v %v", c.FirstName, c.LastName)
	}
}

func (c *User) GetFullName() string {

	if len(c.MiddleName) > 0 {
		return fmt.Sprintf("%v %v %v", c.FirstName, c.MiddleName, c.LastName)
	} else {
		return fmt.Sprintf("%v %v", c.FirstName, c.LastName)
	}
}

func (c *User) TableName() string {
	return "api_user.user"
}

func (c *User) GetEmail() string {
	email := "Not loaded"
	if c.AuthSocial != nil {
		s := *(c.AuthSocial)
		return s.Email
	}
	if c.AuthEmail != nil {
		s := *(c.AuthEmail)
		return s.Email
	}
	return email
}
func (c *User) GetDisplayname() string {
	if ((*c).Profile) != nil {
		p := *((*c).Profile)
		if p.DisplaynameMode == "fullname" {
			return c.GetFullName()
		}
		if p.DisplaynameMode == "abbreviated" {
			return c.GetAbbreviatedName()
		}
		if p.DisplaynameMode == "firstnameonly" {
			return c.FirstName
		}
		if p.DisplaynameMode == "nickname" {
			return p.Name
		}
	}
	return "Unset"
}

func (c *User) GetAbbreviatedName() string {

	if len(c.FirstName) > 0 && len(c.MiddleName) > 0 && len(c.LastName) > 0 {
		return fmt.Sprintf("%v %v. %v.", c.FirstName, c.MiddleName[0:1], c.LastName[0:1])
	} else if len(c.FirstName) > 0 && len(c.LastName) > 0 {
		return fmt.Sprintf("%v %v.", c.FirstName, c.LastName[0:1])
	} else if len(c.FirstName) > 0 {
		return fmt.Sprintf("%v", c.FirstName)
	} else if len(c.LastName) > 0 {
		return fmt.Sprintf("%v", c.LastName)
	} else {
		return ("Undefined")
	}
}

func (c *User) IsNameSet() bool {
	return (len(c.FirstName) + len(c.MiddleName) + len(c.LastName)) > 0
}
func (u *User) RetreiveDetails() ContactDetailsLightData {
	res := ContactDetailsLightData{}
	// if u.AuthStrategy == "MAGICLINK" {
	// 	res.Emails = []string{u.AuthMagiclink.Email}
	// } else if u.AuthStrategy == "EMAILPASSWORD" {
	// 	res.Emails = []string{u.AuthMagiclink.Email}
	// } else {
	// 	res.Emails = []string{u.AuthMagiclink.Email}
	// }
	// res.Phones = []string{u.AuthPhone.Phone}
	return res
}

func (u *User) RetreiveEmail() string {
	// if u.AuthStrategy == "MAGICLINK" {
	// 	return u.AuthMagiclink.Email
	// } else if u.AuthStrategy == "EMAILPASSWORD" {
	// 	return u.AuthEmail.Email
	// } else {
	// 	return u.AuthEmail.Email
	// }
	return u.GetEmail()
}
func (u *User) RetreivePhoneBase() string {
	if u.AuthPhone.ID > 0 {
		return u.AuthPhone.Phone
	}
	return ""
}
func (u *User) RetreiveFullPhone() string {
	if u.AuthPhone.ID > 0 {
		return u.GetUserPhone()
	}
	return ""
}

func (u *User) GetWhereDetails(key string) string {
	email := u.RetreiveEmail()
	phone := u.RetreiveFullPhone()
	if len(email) > 0 && len(phone) > 0 {
		return fmt.Sprintf("( ( %v->'e')::jsonb ? '%s'  OR ( %v->'p')::jsonb ? '%s' ) ", key, email, key, phone)
	} else if len(email) > 0 {
		return fmt.Sprintf("( ( %v->'e')::jsonb ? '%s'   ) ", key, email)
	} else if len(phone) > 0 {
		return fmt.Sprintf("( ( %v->'p')::jsonb ? '%s'   ) ", key, phone)
	} else {
		return " true "
	}

}
func (u *User) GetUserPhone() string {
	if u.AuthPhone.ID > 0 {
		return fmt.Sprintf("%s%s", u.AuthPhone.CountryCode, u.AuthPhone.Phone)
	}

	return ""
}
