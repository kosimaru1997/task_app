package common

import (
	"strings"
	"time"
)

type OnlyDate struct {
	Date time.Time
}

func (onlyDate *OnlyDate) UnmarshalJSON(b []byte) error {
	// Trim surrounding quotes from JSON string
	s := strings.Trim(string(b), "\"")

	// Parse time string using the specified format
	parsedTime, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	// Set the parsed time to the time.Time pointer
	*&onlyDate.Date = parsedTime

	return nil
}
