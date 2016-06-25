package core

import "time"

type Date time.Time

type Time time.Time

type LongTime time.Time

func (t Date) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(t).Format("2006-01-02") + `"`), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(t).Format("2006-01-02 15:04") + `"`), nil
}

func (t LongTime) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(t).Format("2006-01-02 15:04:05") + `"`), nil
}
