package models

import (
	"time"

	"stretches-common-api/utils"

	"gorm.io/gorm"
)

type AuthSocial struct {
	gorm.Model
	ID                uint32 `json:"id"`
	Name              string `json:"name"`
	FirstName         string `json:"first_name"`
	MiddleName        string `json:"middle_name"`
	Strategy          string `json:"strategy"`
	LastName          string `json:"last_name"`
	IdentityToken     string `json:"identity_token"`
	AuthorizationCode string `json:"authorization_code"`

	AvatarUrl    string `json:"avatar_url"`
	Email        string `json:"email"`
	UserId       uint32 `json:"user_id"`
	SocialId     string `json:"apple_id"`
	IsAssociated bool   `json:"is_associated"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	User *User `gorm:"foreignKey:user_id" json:"user,omitempty"`

	SocialConfirmRequests []LinkSocialConfirmRequest `gorm:"foreignKey:social_id" json:"linksocialconfirm,omitempty"`

}

func (c *AuthSocial) GetId() uint32 {
	return c.ID
}
func (c *AuthSocial) GetName() string {
	name := ""
	if len(c.Name) > 0 {
		name = c.Name
	} else if len(c.FirstName) > 0 {
		name = utils.GetName(c.FirstName, c.MiddleName, c.LastName)
	}
	return name
}

func (c *AuthSocial) TableName() string {
	return "api_user.auth_social"
}
