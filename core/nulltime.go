package core

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
    "reflect"
    "time"
)


// NullDate
//------------------------------------------------------------------------------------------------------------------

// Time is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
type NullDate struct {
    Time  time.Time
    Valid bool
}

// Scan implements the Scanner interface.
func (t *NullDate) Scan(value interface{}) error {
    var err error
    switch x := value.(type) {
    case time.Time:
        t.Time = x
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("null: cannot scan type %T into NullDate: %v", value, value)
    }
    t.Valid = err == nil
    return err
}

// Value implements the driver Valuer interface.
func (t NullDate) Value() (driver.Value, error) {
    if !t.Valid {
        return nil, nil
    }
    return t.Time, nil
}

// NewNullDate creates a new Time.
func NewNullDate(t time.Time, valid bool) NullDate {
    return NullDate{
        Time:  t,
        Valid: valid,
    }
}

// TimeFrom creates a new Time that will always be valid.
func NullDateFrom(t time.Time) NullDate {
    return NewNullDate(t, true)
}

// TimeFromPtr creates a new Time that will be null if t is nil.
func NullDateFromPtr(t *time.Time) NullDate {
    if t == nil {
        return NewNullDate(time.Time{}, false)
    }
    return NewNullDate(*t, true)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t NullDate) MarshalJSON() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    //return t.Time.MarshalJSON()
    return []byte(`"` + t.Time.Format("2006-01-02") + `"`), nil

}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullDate and friends)
// and null input.
func (t *NullDate) UnmarshalJSON(data []byte) error {
    var err error
    var v interface{}
    if err = json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch x := v.(type) {
    case string:
        err = t.Time.UnmarshalJSON(data)
    case map[string]interface{}:
        ti, tiOK := x["Time"].(string)
        valid, validOK := x["Valid"].(bool)
        if !tiOK || !validOK {
            return fmt.Errorf(`json: unmarshalling object into Go value of type NullDate requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
        }
        err = t.Time.UnmarshalText([]byte(ti))
        t.Valid = valid
        return err
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
    }
    t.Valid = err == nil
    return err
}

func (t NullDate) MarshalText() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    return t.Time.MarshalText()
}

func (t *NullDate) UnmarshalText(text []byte) error {
    str := string(text)
    if str == "" || str == "null" {
        t.Valid = false
        return nil
    }
    if err := t.Time.UnmarshalText(text); err != nil {
        return err
    }
    t.Valid = true
    return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *NullDate) SetValid(v time.Time) {
    t.Time = v
    t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t NullDate) Ptr() *time.Time {
    if !t.Valid {
        return nil
    }
    return &t.Time
}

// NullTime
//------------------------------------------------------------------------------------------------------------------

// Time is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
type NullTime struct {
    Time  time.Time
    Valid bool
}

// Scan implements the Scanner interface.
func (t *NullTime) Scan(value interface{}) error {
    var err error
    switch x := value.(type) {
    case time.Time:
        t.Time = x
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("null: cannot scan type %T into NullTime: %v", value, value)
    }
    t.Valid = err == nil
    return err
}

// Value implements the driver Valuer interface.
func (t NullTime) Value() (driver.Value, error) {
    if !t.Valid {
        return nil, nil
    }
    return t.Time, nil
}

// NewNullTime creates a new Time.
func NewNullTime(t time.Time, valid bool) NullTime {
    return NullTime{
        Time:  t,
        Valid: valid,
    }
}

// TimeFrom creates a new Time that will always be valid.
func NullTimeFrom(t time.Time) NullTime {
    return NewNullTime(t, true)
}

// TimeFromPtr creates a new Time that will be null if t is nil.
func NullTimeFromPtr(t *time.Time) NullTime {
    if t == nil {
        return NewNullTime(time.Time{}, false)
    }
    return NewNullTime(*t, true)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t NullTime) MarshalJSON() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    //return t.Time.MarshalJSON()
    return []byte(`"` + t.Time.Format("2006-01-02 15:04") + `"`), nil

}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (t *NullTime) UnmarshalJSON(data []byte) error {
    var err error
    var v interface{}
    if err = json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch x := v.(type) {
    case string:
        err = t.Time.UnmarshalJSON(data)
    case map[string]interface{}:
        ti, tiOK := x["Time"].(string)
        valid, validOK := x["Valid"].(bool)
        if !tiOK || !validOK {
            return fmt.Errorf(`json: unmarshalling object into Go value of type NullTime requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
        }
        err = t.Time.UnmarshalText([]byte(ti))
        t.Valid = valid
        return err
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
    }
    t.Valid = err == nil
    return err
}

func (t NullTime) MarshalText() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    return t.Time.MarshalText()
}

func (t *NullTime) UnmarshalText(text []byte) error {
    str := string(text)
    if str == "" || str == "null" {
        t.Valid = false
        return nil
    }
    if err := t.Time.UnmarshalText(text); err != nil {
        return err
    }
    t.Valid = true
    return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *NullTime) SetValid(v time.Time) {
    t.Time = v
    t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t NullTime) Ptr() *time.Time {
    if !t.Valid {
        return nil
    }
    return &t.Time
}

// NullLongTime
//------------------------------------------------------------------------------------------------------------------

// Time is a nullable time.Time. It supports SQL and JSON serialization.
// It will marshal to null if null.
type NullLongTime struct {
    Time  time.Time
    Valid bool
}

// Scan implements the Scanner interface.
func (t *NullLongTime) Scan(value interface{}) error {
    var err error
    switch x := value.(type) {
    case time.Time:
        t.Time = x
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("null: cannot scan type %T into NullLongTime: %v", value, value)
    }
    t.Valid = err == nil
    return err
}

// Value implements the driver Valuer interface.
func (t NullLongTime) Value() (driver.Value, error) {
    if !t.Valid {
        return nil, nil
    }
    return t.Time, nil
}

// NewNullLongTime creates a new Time.
func NewNullLongTime(t time.Time, valid bool) NullLongTime {
    return NullLongTime{
        Time:  t,
        Valid: valid,
    }
}

// TimeFrom creates a new Time that will always be valid.
func NullLongTimeFrom(t time.Time) NullLongTime {
    return NewNullLongTime(t, true)
}

// TimeFromPtr creates a new Time that will be null if t is nil.
func NullLongTimeFromPtr(t *time.Time) NullLongTime {
    if t == nil {
        return NewNullLongTime(time.Time{}, false)
    }
    return NewNullLongTime(*t, true)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this time is null.
func (t NullLongTime) MarshalJSON() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    //return t.Time.MarshalJSON()
    return []byte(`"` + t.Time.Format("2006-01-02 15:04:05") + `"`), nil

}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullLongTime and friends)
// and null input.
func (t *NullLongTime) UnmarshalJSON(data []byte) error {
    var err error
    var v interface{}
    if err = json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch x := v.(type) {
    case string:
        err = t.Time.UnmarshalJSON(data)
    case map[string]interface{}:
        ti, tiOK := x["Time"].(string)
        valid, validOK := x["Valid"].(bool)
        if !tiOK || !validOK {
            return fmt.Errorf(`json: unmarshalling object into Go value of type NullLongTime requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
        }
        err = t.Time.UnmarshalText([]byte(ti))
        t.Valid = valid
        return err
    case nil:
        t.Valid = false
        return nil
    default:
        err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.Time", reflect.TypeOf(v).Name())
    }
    t.Valid = err == nil
    return err
}

func (t NullLongTime) MarshalText() ([]byte, error) {
    if !t.Valid {
        return []byte("null"), nil
    }
    return t.Time.MarshalText()
}

func (t *NullLongTime) UnmarshalText(text []byte) error {
    str := string(text)
    if str == "" || str == "null" {
        t.Valid = false
        return nil
    }
    if err := t.Time.UnmarshalText(text); err != nil {
        return err
    }
    t.Valid = true
    return nil
}

// SetValid changes this Time's value and sets it to be non-null.
func (t *NullLongTime) SetValid(v time.Time) {
    t.Time = v
    t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t NullLongTime) Ptr() *time.Time {
    if !t.Valid {
        return nil
    }
    return &t.Time
}
