package formatters

import (
	"fmt"
	m "stretches-common-api/models"
	"time"
)

type DeviceFormattedItem struct {
	CreatedAt          *time.Time `json:"created_at"`
	DeviceId           string     `json:"device_id"`
	DeviceModel        string     `json:"device_model"`
	DeviceManufacturer string     `json:"device_manufacturer"`
	DeviceName         string     `json:"device_name"`
	DeviceBrand        string     `json:"device_brand"`
	OsName             string     `json:"os_"`
	OsVersion          string     `json:"os_version"`
	Extra              string     `json:"extra"`
	Id                 int        `json:"id"`
}

func FormatDevices(items []m.Device) []DeviceFormattedItem {
	res := []DeviceFormattedItem{}
	for _, v := range items {
		res = append(res, FormatDevice(v))
	}
	return res
}
func FormatLinkDeviceUsers(items []m.LinkDeviceUser) []DeviceFormattedItem {
	res := []DeviceFormattedItem{}
	for _, v := range items {
		if v.Device != nil {
			res = append(res, FormatDevice(*(v.Device)))
		}
	}
	return res
}
func FormatDevice(v m.Device) DeviceFormattedItem {
	res := DeviceFormattedItem{}
	res = DeviceFormattedItem{
		CreatedAt:          v.CreatedAt,
		DeviceId:           v.Identifier,
		DeviceModel:        v.Model_,
		DeviceBrand:        v.Brand,
		OsName:             v.OsName,
		OsVersion:          v.OsVersion,
		Extra:              v.Extra,
		DeviceManufacturer: v.Manufacturer,
		DeviceName:         v.Name,
	}
	return res
}

type AuthstrategyFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	Name      string     `json:"name"`
	Id        int        `json:"id"`
}

func FormatAuthstrategies(items []m.AuthStrategy) []AuthstrategyFormattedItem {
	res := []AuthstrategyFormattedItem{}
	for _, v := range items {
		res = append(res, FormatAuthstrategy(v))
	}
	return res
}

func FormatAuthstrategy(v m.AuthStrategy) AuthstrategyFormattedItem {
	res := AuthstrategyFormattedItem{}
	res = AuthstrategyFormattedItem{
		CreatedAt: v.CreatedAt,
	}
	if v.AuthAppleId != 0 {
		res.Name = "Apple"
	}
	if v.AuthGoogleId != 0 {
		res.Name = "Google"
	}
	if v.AuthFacebookId != 0 {
		res.Name = "Facebook"
	}

	if v.AuthEmailpasswordId != 0 {
		res.Name = "Email/Password"
	}
	return res
}

type LoginattemptsFormattedItem struct {
	CreatedAt *time.Time `json:"created_at"`
	Id        int        `json:"id"`
}

func FormatLoginattempts(items []m.Loginhistory) []LoginattemptsFormattedItem {
	res := []LoginattemptsFormattedItem{}
	for _, v := range items {
		res = append(res, FormatLoginattempt(v))
	}
	return res
}

func FormatLoginattempt(v m.Loginhistory) LoginattemptsFormattedItem {
	res := LoginattemptsFormattedItem{}
	res = LoginattemptsFormattedItem{
		CreatedAt: v.CreatedAt,
	}
	return res
}

type RequestFormattedItem struct {
	CreatedAt   *time.Time `json:"created_at"`
	Type        string     `json:"type"`
	Link        string     `json:"link"`
	Dest        string     `json:"dest"`
	Token       string     `json:"token"`
	IsConfirmed bool       `json:"is_confirmed"`
	IsSent      bool       `json:"is_sent"`
}

func FormatRequests(u m.User /*items []m.EmailConfirmRequest, device []m.DeviceConfirmRequest, socials []m.LinkSocialConfirmRequest*/) []RequestFormattedItem {
	res := []RequestFormattedItem{}
	if u.AuthEmail != nil {
		for _, v := range (*u.AuthEmail).EmailConfirmRequests {
			res = append(res, FormatRequest(v))
		}
	}

	for _, v := range u.Devices {
		for _, v := range v.DeviceConfirmRequests {
			res = append(res, FormatRequest2(v))
		}
	}
	if u.AuthSocial != nil {
		for _, v := range (*u.AuthSocial).SocialConfirmRequests {
			res = append(res, FormatRequest3(v))
		}
	}
	return res
}

func FormatRequest(v m.EmailConfirmRequest) RequestFormattedItem {
	res := RequestFormattedItem{}
	res = RequestFormattedItem{
		IsSent:    v.Sent,
		CreatedAt: v.CreatedAt,
		Dest:      v.Email,
		Token:     v.Token,
		Link:      fmt.Sprintf("https://stretches.weand.co.uk/user/live/v1/auth/emailpassword/confirm/%v/%v", v.Email, v.Token), // v.Link,
		//		Link:        v.Link,
		IsConfirmed: v.Confirmed,
		Type:        "EmailConfirm",
	}
	return res
}

func FormatRequest2(v m.DeviceConfirmRequest) RequestFormattedItem {
	res := RequestFormattedItem{}
	res = RequestFormattedItem{
		CreatedAt:   v.CreatedAt,
		IsConfirmed: v.Confirmed,
		// Link:        v.Link,
		Link:   fmt.Sprintf("https://stretches.weand.co.uk/user/live/v1/auth/device/confirm/%v/%v", v.Email, v.Token), // v.Link,
		Token:  v.Token,
		IsSent: v.Sent,
		Dest:   v.DeviceId,
		Type:   "DeviceConfirm",
	}
	return res
}

func FormatRequest3(v m.LinkSocialConfirmRequest) RequestFormattedItem {
	res := RequestFormattedItem{}
	res = RequestFormattedItem{
		CreatedAt:   v.CreatedAt,
		IsSent:      v.Sent,
		Link:        fmt.Sprintf("https://stretches.weand.co.uk/user/live/v1/auth/linksocial/confirm/%v/%v", v.Email, v.Token), // v.Link,
		IsConfirmed: v.Confirmed,
		Token:       v.Token,
		Type:        "LinkSocialConfirm",
	}
	return res
}
