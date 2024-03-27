package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Interests struct {
	Sendmoney int `json:"SENDMONEY"`
	Moneypool int `json:"MONEYPOOL"`
	Splitbill int `json:"SPLITBILL"`
	Forex     int `json:"FOREX"`
	Crypto    int `json:"CRYPTO"`
	Invest    int `json:"INVEST"`
	Analyze   int `json:"ANALYZE"`
	Share     int `json:"SHARE"`
	Budget    int `json:"BUDGET"`
	Deals     int `json:"DEALS"`
	QR        int `json:"QR"`

	Family   int `json:"FAMILY"`
	Cashback int `json:"CASHBACK"`
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *Interests) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Interests{}
	err := json.Unmarshal(bytes, &result)
	*j = Interests(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j Interests) Value() (driver.Value, error) {
	return json.Marshal(j)
}
