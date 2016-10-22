package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type NullInt struct {
	sql.NullInt64
}

func NewNullInt(i int64, valid bool) NullInt {
	return NullInt{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: valid,
		},
	}
}

func NullIntFrom(i int64) NullInt {
	return NewNullInt(i, true)
}

func NullIntFromPtr(i *int64) NullInt {
	if i == nil {
		return NewNullInt(0, false)
	}
	return NewNullInt(*i, true)
}

func (i *NullInt) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.(type) {
	case float64:
		// Unmarshal again, directly to int64, to avoid intermediate float64
		err = json.Unmarshal(data, &i.Int64)
	case map[string]interface{}:
		err = json.Unmarshal(data, &i.NullInt64)
	case string:
		i.UnmarshalText(data)
	case nil:
		i.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullInt", reflect.TypeOf(v).Name())
	}
	i.Valid = err == nil
	return err
}

func (i *NullInt) UnmarshalText(text []byte) error {
	str := strings.Trim(string(text), `"`)
	if str == "" || str == "null" {
		i.Valid = false
		return nil
	}
	var err error
	i.Int64, err = strconv.ParseInt(str, 10, 64)
	i.Valid = err == nil

	return err
}

func (i NullInt) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(i.Int64, 10)), nil
}

func (i NullInt) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(i.Int64, 10)), nil
}

func (i *NullInt) SetValid(n int64) {
	i.Int64 = n
	i.Valid = true
}

func (i NullInt) Ptr() *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

func (i NullInt) IsZero() bool {
	return !i.Valid
}
