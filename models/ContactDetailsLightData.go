package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ContactDetailsLightData struct {
	Emails []string `json:"e,omitempty"`
	Phones []string `json:"p,omitempty"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ContactDetailsLightData) Reformat() {
	for i := 0; i < len(j.Emails); i++ {
		j.Emails[i] = strings.ToLower(j.Emails[i])
	}
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ContactDetailsLightData) ToStr() string {
	res := ""
	for i := 0; i < len(j.Emails); i++ {
		res = fmt.Sprintf("%s %s", res, j.Emails[i])
	}
	for i := 0; i < len(j.Phones); i++ {
		res = fmt.Sprintf("%s %s", res, j.Phones[i])
	}
	return res
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ContactDetailsLightData) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := ContactDetailsLightData{}
	err := json.Unmarshal(bytes, &result)
	*j = ContactDetailsLightData(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j ContactDetailsLightData) Value() (driver.Value, error) {
	return json.Marshal(j)
}
