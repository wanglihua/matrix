package core

import (
	"time"
	"strings"
)

type Date time.Time

type Time time.Time

type LongTime time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02") + `"`), nil
}

func (t *Date) UnmarshalJSON(data []byte) error {
	parsed_time, err := time.Parse("2006-01-02", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	*t = Date(parsed_time)
	return nil
}

func (t Date) String() string {
	return time.Time(t).Format("2006-01-02")
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04") + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	parsed_time, err := time.Parse("2006-01-02 15:04", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	*t = Time(parsed_time)
	return nil
}

func (t Time) String() string {
	return time.Time(t).Format("2006-01-02 15:04")
}

func (t LongTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04:05") + `"`), nil
}

func (t *LongTime) UnmarshalJSON(data []byte) error {
	parsed_time, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	*t = LongTime(parsed_time)
	return nil
}

func (t LongTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}
