package models

import (
	"fmt"
	"stretches-common-api/structs"
	"stretches-common-api/timezone"
	util "stretches-common-api/utils"
	"strings"
	"time"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/teambition/rrule-go"
	"golang.org/x/text/language"
	"gorm.io/gorm"

	"github.com/skillcoder/hrrule-go"
)

type Event struct {
	gorm.Model
	ID        uint32     `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name        string `json:"name"`
	Rrule       string `json:"rrule"`
	IsRecurring bool   `json:"is_recurring"`
	IsAllDay    bool   `json:"is_all_day"`
	Duration    int32  `json:"duration"`

	StartdateEnteredOffset   *int       `json:"startdate_entered_offset,omitempty"`
	StartdateEnteredTimezone *string    `json:"startdate_entered_timezone,omitempty"`
	StartdateEnteredDatetime *time.Time `json:"startdate_entered_datetime,omitempty"`

	StartdateUtc    *time.Time `json:"startdate_utc,omitempty"`
	EndrangedateUtc *time.Time `json:"endrangedate_utc,omitempty"`

	Type         string `json:"type"`
	RecurringKey string `json:"recurringkey"`

	//POPULATED FIELDS
	Exceptions []Eventexception  `gorm:"foreignKey:event_id" json:"exceptions,omitempty"`
	LinkTeams  []LinkEventTeam   `gorm:"foreignKey:event_id" json:"linkteams,omitempty"`
	Replies    []LinkEventPerson `gorm:"foreignKey:event_id" json:"replies,omitempty"`

	Batches       []SendBatch         `gorm:"foreignKey:related_id" json:"batches,omitempty"`
	LinkLocations []LinkEventLocation `gorm:"foreignKey:EventId;references:ID" json:"linklocations,omitempty"`
}

func (c *Event) GetId() uint32 {
	return c.ID
}

func (c *Event) TableName() string {
	return "api_event.event"
}

func (c *Event) GetNameOrGenericName() string {
	if len(c.Name) > 0 && c.Name != "?" {
		return c.Name
	} else {
		switch c.Type {
		case "practice":
			return "Practice"
		case "staff":
			return "Staff meeting"
		case "rehearsal":
			return "Rehearsal"
		}
	}
	return "Event"
}
func (c *Event) GetRRuleDesc() (bool, string) {
	isReadyForRecurring, _ := c.GetRRuleWithDate()
	if !isReadyForRecurring {
		return true, ""
	} else {

		hRule, err := hrrule.New(i18n.NewBundle(language.AmericanEnglish))
		if err != nil {
			return false, ""
		}
		rrule := strings.Replace(c.Rrule, "RRULE:", "", 1)
		rOption, err := hrrule.StrToROption(rrule)
		if err != nil {
			return false, ""
		}

		nlString, err := hRule.Humanize(rOption, "en-US")
		if err != nil {
			return false, ""
		}
		return true, nlString

	}
}
func (c *Event) GetRRuleWithDate() (bool, string) {
	if c.StartdateUtc == nil {
		return false, ""
	}
	if !c.IsRecurring {
		return false, ""
	} else {
		startStr := c.StartdateUtc.UTC().Format(timezone.DatetimeCompactFormat)
		return true, fmt.Sprintf("DTSTART:%v\n%v", startStr, c.Rrule)
	}
}
func (c *Event) GetOccurences() (bool, []time.Time) {
	isReadyForRecurring, ruleWithDate := c.GetRRuleWithDate()
	if !isReadyForRecurring {
		return true, []time.Time{}
	} else {
		fmt.Printf("%v\n", ruleWithDate)

		s, err := rrule.StrToRRuleSet(ruleWithDate)

		if err != nil {
			fmt.Printf("%v\n", err)
			return false, []time.Time{}
		} else {
			return true, s.All()
		}

	}
}
func (c *Event) CanEdit(me structs.Me) bool {
	for _, liTeam := range c.LinkTeams {
		if liTeam.Team != nil {
			if liTeam.Team.CanEdit(me) {
				return true
			}
		}
	}
	return false
}

func (c *Event) GetPastOrFuture() string {
	now := time.Now()
	if c.HasStart() {
		return "UNSCHEDULED"
	} else {
		if c.StartdateUtc == nil {
			return "UNSCHEDULED"
		}
		d := (*c.StartdateUtc)
		if d.After(now) {
			return "FUTURE"
		} else {
			if c.Duration > 0 {
				end := d.Add(time.Hour * time.Duration(c.Duration))
				if end.After(now) {
					return "NOW"
				} else {
					return "PAST"
				}
			} else {
				if d.Before(now) {
					return "PAST"
				}
			}
		}
	}
	return "UNSET"
}
func (c *Event) HasStart() bool {
	if c.StartdateUtc == nil {
		return false
	}
	threshold := time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)
	isBeforeThreshold := c.StartdateUtc.Before(threshold)
	return isBeforeThreshold
}
func (c *Event) GenerateRecurringKey() string {
	key := util.RandomString(16)
	c.RecurringKey = key
	return key
}
