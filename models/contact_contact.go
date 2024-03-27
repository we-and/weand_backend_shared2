package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"database/sql/driver"
	"time"
	_ "time"
)

type ContactUploadAddressesRequest struct {
	Street   string `json:"st"`
	Postcode string `json:"pc"`
	City     string `json:"ci"`
	Country  string `json:"co"`
	Region   string `json:"re"`
}
type ContactDetailsData struct {
	Emails    []string                        `json:"e,omitempty"`
	Phones    []string                        `json:"p,omitempty"`
	Addresses []ContactUploadAddressesRequest `json:"a,omitempty"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ContactDetailsData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := ContactDetailsData{}
	err := json.Unmarshal(bytes, &result)
	*j = ContactDetailsData(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j ContactDetailsData) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Contact struct {
	gorm.Model
	ID          uint32             `json:"id",gorm:"primaryKey"`
	FirstName   string             `json:"first_name`
	LastName    string             `json:"last_name`
	MiddleName  string             `json:"middle_name`
	Identifier  string             `json:"identifier`
	Birthday    string             `json:"birthday`
	DisplayName string             `json:"display_name`
	CreatedAt   *time.Time         `json:"created_at,omitempty"`
	UpdatedAt   *time.Time         `json:"updated_at,omitempty"`
	DeletedAt   *time.Time         `json:"deleted_at,omitempty"`
	Deleted     bool               `json:"deleted,omitempty"`
	UserId      uint32             `json:"user_id,omitempty"`
	OwnerId     uint32             `json:"owner_id,omitempty"`
	Details     ContactDetailsData `json:"details,omitempty"`
}

func (c *Contact) GetId() uint32 {
	return c.ID
}

func (c *Contact) TableName() string {
	return "api_contact.contact"
}
