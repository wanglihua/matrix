package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

type NullString struct {
	sql.NullString
}

func NullStringFrom(s string) NullString {
	return NewNullString(s, true)
}

func NullStringFromPtr(s *string) NullString {
	if s == nil {
		return NewNullString("", false)
	}
	return NewNullString(*s, true)
}

func NewNullString(s string, valid bool) NullString {
	return NullString{
		NullString: sql.NullString{
			String: s,
			Valid:  valid,
		},
	}
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		s.String = x
	case map[string]interface{}:
		err = json.Unmarshal(data, &s.NullString)
	case nil:
		s.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullString", reflect.TypeOf(v).Name())
	}
	s.Valid = err == nil
	return err
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s NullString) MarshalText() ([]byte, error) {
	if !s.Valid {
		return []byte{}, nil
	}
	return []byte(s.String), nil
}

func (s *NullString) UnmarshalText(text []byte) error {
	s.String = string(text)
	s.Valid = s.String != ""
	return nil
}

func (s *NullString) SetValid(v string) {
	s.String = v
	s.Valid = true
}

func (s NullString) Ptr() *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func (s NullString) IsZero() bool {
	return !s.Valid
}
