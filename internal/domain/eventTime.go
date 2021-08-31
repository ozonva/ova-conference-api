package domain

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type EventTime struct {
	time.Time
}

func (et *EventTime) Scan(value interface{}) error {
	et.Time = value.(time.Time)
	return nil
}

func (et EventTime) Value() (driver.Value, error) {
	return et.Time, nil
}
func (et EventTime) String() string {
	return et.Format(layout)
}

const layout = "2006-01-02 15:04:05"

func (et *EventTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}
	et.Time, err = time.Parse(layout, s)
	return
}

func (et EventTime) MarshalJSON() ([]byte, error) {
	if et.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, et.Time.Format(layout))), nil
}
