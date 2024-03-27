package models

import (
	"fmt"
	"time"
	_ "time"

	"gorm.io/gorm"
)

type Unregistered struct {
	gorm.Model
	ID uint32 `json:"id" gorm:"primaryKey"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`

	//Details structs.UserDetails `json:"details"`
	//Email     string              `json:"email"`
	//PhoneExt  string              `json:"phone_ext"`
	//PhoneBase string              `json:"phone_base"`
}

func (c *Unregistered) GetId() uint32 {
	return c.ID
}

func (c *Unregistered) TableName() string {
	return "api_user.unregistered_user"
}

//func (u *Unregistered) GetPhone() string {
//	return fmt.Sprintf("%s%s", u.PhoneExt, u.PhoneExt)

// }
func (u *Unregistered) GetDisplayname() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)

}
