package core

import (
    "database/sql"
    "encoding/json"
    "errors"
    "fmt"
    "reflect"
)

type NullBool struct {
    sql.NullBool
}

func NewNullBool(b bool, valid bool) NullBool {
    return NullBool{
        NullBool: sql.NullBool{
            Bool:  b,
            Valid: valid,
        },
    }
}

func NullBoolFrom(b bool) NullBool {
    return NewNullBool(b, true)
}

func NullBoolFromPtr(b *bool) NullBool {
    if b == nil {
        return NewNullBool(false, false)
    }
    return NewNullBool(*b, true)
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
    var err error
    var v interface{}
    if err = json.Unmarshal(data, &v); err != nil {
        return err
    }
    switch x := v.(type) {
    case bool:
        b.Bool = x
    case map[string]interface{}:
        err = json.Unmarshal(data, &b.NullBool)
    case nil:
        b.Valid = false
        return nil
    default:
        err = fmt.Errorf("json: cannot unmarshal %v into Go value of type NullBool", reflect.TypeOf(v).Name())
    }
    b.Valid = err == nil
    return err
}

func (b *NullBool) UnmarshalText(text []byte) error {
    str := string(text)
    switch str {
    case "", "null":
        b.Valid = false
        return nil
    case "true":
        b.Bool = true
    case "false":
        b.Bool = false
    default:
        b.Valid = false
        return errors.New("invalid input:" + str)
    }
    b.Valid = true
    return nil
}

func (b NullBool) MarshalJSON() ([]byte, error) {
    if !b.Valid {
        return []byte("null"), nil
    }
    if !b.Bool {
        return []byte("false"), nil
    }
    return []byte("true"), nil
}

func (b NullBool) MarshalText() ([]byte, error) {
    if !b.Valid {
        return []byte{}, nil
    }
    if !b.Bool {
        return []byte("false"), nil
    }
    return []byte("true"), nil
}

func (b *NullBool) SetValid(v bool) {
    b.Bool = v
    b.Valid = true
}

func (b NullBool) Ptr() *bool {
    if !b.Valid {
        return nil
    }
    return &b.Bool
}

func (b NullBool) IsZero() bool {
    return !b.Valid
}
