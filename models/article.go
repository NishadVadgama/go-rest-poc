package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Article struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        Tags   `json:"tags"`
}

type Tags []string

func (s *Tags) Scan(value interface{}) error {
	// Tags can be null
	if value == nil {
		*s = []string{}
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}

func (s Tags) Value() (driver.Value, error) {
	return json.Marshal(s)
}
