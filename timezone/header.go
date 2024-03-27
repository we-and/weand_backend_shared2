package timezone

import (
	"fmt"
	"math"
	"time"
)

const (
	DatetimeWithTzFormat       = "2006-01-02T15:04:05Z"
	DatetimeWithOffsetFormat   = "2006-01-02T15:04:05-0700"
	DatetimeDbqueryFormat      = "2006-01-02 15:04:05-07"
	DateFormat                 = "2006-01-02"
	TimeFormat                 = "15:04"
	TimeWithOffsetFormat       = "15:04-0700"
	OffsetFormat               = "-0700"
	DatetimeFormat             = "2006-01-02T15:04:05"
	HumanDatetimeFormat        = "Monday 2th January 15:04"
	HumanTimeFormat            = "15:04"
	HumanDateFormat            = "Monday 2th January"
	DatetimeSpacedSimpleFormat = "2006-01-02 15:04"
	DatetimeCompactFormat      = "20060102T150405Z"
)

type TzData struct {
	Name   string
	Offset string
	Loc    *time.Location
}

func OffsetStrToMinutes(offset string) (error, int) {
	offsetStandard := "2000-01-01T00:00:00"
	fullOffset := fmt.Sprintf("%v%v", offsetStandard, offset)
	starttimeOffsetDate, err := time.Parse(DatetimeWithOffsetFormat, fullOffset)
	if err != nil {
		//		app.SaveBadRequestAndReturn(r,  "Request fields incorrect. Expecting offset.", errors.New(""), "MS001-003")
		return err, 0
	}
	return nil, TimeToIntOffset(starttimeOffsetDate)

}
func TimeToIntOffset(t time.Time) int {
	zeroDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	diffDuration := zeroDate.Sub(t)

	return int(math.Round(diffDuration.Minutes()))

}
