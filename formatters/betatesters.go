package formatters

import (
	models "stretches-common-api/models"
	"stretches-common-api/publicid"
	"time"
)

type BetatesterFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	Pid       uint32     `json:"pid"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
}

func FormatBetatesters(vv []models.Betatester) []BetatesterFormattedItem {
	res := []BetatesterFormattedItem{}
	for _, v := range vv {
		res = append(res, FormatBetatester(v))
	}
	return res
}
func FormatBetatester(v models.Betatester) BetatesterFormattedItem {

	res := BetatesterFormattedItem{
		Pid:   publicid.Obfuscate32bit(v.ID),
		Name:  v.Name,
		Email: v.Email,
		Phone: v.Phone,
	}
	return res
}
