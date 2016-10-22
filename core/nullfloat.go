package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NullFloat struct {
	sql.NullFloat64
}

func NewNullFloat(f float64, valid bool) NullFloat {
	return NullFloat{
		NullFloat64: sql.NullFloat64{
			Float64: f,
			Valid:   valid,
		},
	}
}

func NullFloatFrom(f float64) NullFloat {
	return NewNullFloat(f, true)
}

func NullFloatFromPtr(f *float64) NullFloat {
	if f == nil {
		return NewNullFloat(0, false)
	}
	return NewNullFloat(*f, true)
}

func (f *NullFloat) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case float64:
		f.Float64 = float64(x)
	case map[string]interface{}:
		err = json.Unmarshal(data, &f.NullFloat64)
	case string:
		f.UnmarshalText(data)
	case nil:
		f.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullFloat", reflect.TypeOf(v).Name())
	}
	f.Valid = err == nil
	return err
}

func (f *NullFloat) UnmarshalText(text []byte) error {
	str := strings.Trim(string(text), `"`)
	if str == "" || str == "null" {
		f.Valid = false
		return nil
	}
	var err error
	f.Float64, err = strconv.ParseFloat(string(text), 64)
	f.Valid = err == nil
	return err
}

func (f NullFloat) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatFloat(f.Float64, 'f', -1, 64)), nil
}

func (f NullFloat) MarshalText() ([]byte, error) {
	if !f.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatFloat(f.Float64, 'f', -1, 64)), nil
}

func (f *NullFloat) SetValid(n float64) {
	f.Float64 = n
	f.Valid = true
}

func (f NullFloat) Ptr() *float64 {
	if !f.Valid {
		return nil
	}
	return &f.Float64
}

func (f NullFloat) IsZero() bool {
	return !f.Valid
}
