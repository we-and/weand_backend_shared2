package structs

import(
	"errors"
	"fmt"

	"database/sql/driver"
	"encoding/json"
)

type EmailData struct {
	Email string `json:"email"`
	Type  string `json:"type"`

}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *EmailData) Scan(value interface{}) error {
  bytes, ok := value.([]byte)
  if !ok {
    return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
  }

  result := EmailData{}
  err := json.Unmarshal(bytes, &result)
  *j = EmailData(result) 
	return err
}

// Value return json value, implement driver.Valuer interface
func (j EmailData) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type PhoneData struct {
	Base string `json:"base"`
	Ext  string `json:"ext"`
	Iso  string `json:"iso"`
	Type string `json:"type"`
}
func (j *PhoneData) Build() string {
	return fmt.Sprintf("%v%v",j.Ext,j.Base)
}
// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *PhoneData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
	  return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
  
	result := PhoneData{}
	err := json.Unmarshal(bytes, &result)
	*j = PhoneData(result)	
	return err
  }
  
  // Value return json value, implement driver.Valuer interface
  func (j PhoneData) Value() (driver.Value, error) {
	return json.Marshal(j)
  }
type UserDetails struct {
	Emails []EmailData `json:"emails"`
	Phones []PhoneData `json:"phones"`
}


// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *UserDetails) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
	  return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
  
	result := UserDetails{}
	err := json.Unmarshal(bytes, &result)
	*j = UserDetails(result)
	return err
  }
  
  // Value return json value, implement driver.Valuer interface
  func (j UserDetails) Value() (driver.Value, error) {
	return json.Marshal(j)
  }