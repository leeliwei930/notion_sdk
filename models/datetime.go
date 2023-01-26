package models

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

func (dp *DateTime) UnmarshalJSON(dateTime []byte) error {
	dateOnly := "2006-01-02"
	var dateTimeString string
	json.Unmarshal(dateTime, &dateTimeString)
	// convert to datetime format
	parsedDateTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err == nil {
		dp.Time = parsedDateTime
		return nil
	}

	// fallback for convert date only format
	date, dateErr := time.Parse(dateOnly, dateTimeString)
	if dateErr != nil {
		return dateErr
	}
	dp.Time = date
	return nil
}

func (dp *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dp.String())
}
