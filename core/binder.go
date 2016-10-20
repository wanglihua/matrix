package core

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/revel/revel"
)

var (
	DateBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			r, err := time.ParseInLocation("2006-01-02", val, time.Local)
			if err == nil {
				return reflect.ValueOf(Date(r))
			}
			parseError := errors.New(val + "不是有效的日期格式！")
			if strings.TrimSpace(val) == "" {
				parseError = errors.New("日期不能为空字符串！")
			}
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(Date)
			output[name] = time.Time(t).Format("2006-01-02")
		},
	}

	TimeBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			r, err := time.ParseInLocation("2006-01-02 15:04", val, time.Local)
			if err == nil {
				return reflect.ValueOf(Time(r))
			}
			parseError := errors.New(val + "不是有效的时间格式！")
			if strings.TrimSpace(val) == "" {
				parseError = errors.New("时间不能为空字符串！")
			}
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(Time)
			output[name] = time.Time(t).Format("2006-01-02 15:04")
		},
	}

	LongTimeBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			r, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
			if err == nil {
				return reflect.ValueOf(LongTime(r))
			}
			parseError := errors.New(val + "不是有效的时间格式！")
			if strings.TrimSpace(val) == "" {
				parseError = errors.New("时间不能为空字符串！")
			}
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(LongTime)
			output[name] = time.Time(t).Format("2006-01-02 15:04:05")
		},
	}

	NullDateBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullDate(time.Time{}, false))
			}
			r, err := time.ParseInLocation("2006-01-02", val, time.Local)
			if err == nil {
				return reflect.ValueOf(NewNullDate(r, true))
			}
			parseError := errors.New(val + "不是有效的日期格式！")
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullDate)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = t.Time.Format("2006-01-02")
			}
		},
	}

	NullTimeBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullTime(time.Time{}, false))
			}
			r, err := time.ParseInLocation("2006-01-02 15:04", val, time.Local)
			if err == nil {
				return reflect.ValueOf(NewNullTime(r, true))
			}
			parseError := errors.New(val + "不是有效的时间格式！")
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullTime)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = t.Time.Format("2006-01-02 15:04")
			}
		},
	}

	NullLongTimeBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullLongTime(time.Time{}, false))
			}
			r, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local)
			if err == nil {
				return reflect.ValueOf(NewNullLongTime(r, true))
			}
			parseError := errors.New(val + "不是有效的时间格式！")
			HandleError(parseError)
			return reflect.Zero(typ)
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullLongTime)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = t.Time.Format("2006-01-02 15:04:05")
			}
		},
	}

	// Booleans support a couple different value formats:
	// "true" and "false"
	// "on" and "" (a checkbox)
	// "1" and "0" (why not)
	NullBoolBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			v := strings.TrimSpace(strings.ToLower(val))
			if v == "" {
				return reflect.ValueOf(NewNullBool(false, false))
			}
			switch v {
			case "true", "on", "1":
				return reflect.ValueOf(NewNullBool(true, true))
			case "false", "off", "0":
				return reflect.ValueOf(NewNullBool(false, true))
			}
			parseError := errors.New(val + "不是有效的布尔值格式！")
			HandleError(parseError)
			return reflect.ValueOf(NewNullBool(false, false))
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullBool)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = fmt.Sprintf("%t", t.Bool)
			}
		},
	}

	NullIntBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullInt(0, false))
			}
			intValue, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				HandleError(err)
				return reflect.ValueOf(NewNullInt(0, false))
			}
			return reflect.ValueOf(NewNullInt(intValue, true))
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullInt)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = fmt.Sprintf("%d", t.Int64)
			}
		},
	}

	NullFloatBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullFloat(0.0, false))
			}
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				HandleError(err)
				return reflect.ValueOf(NewNullFloat(0.0, false))
			}
			return reflect.ValueOf(NewNullFloat(floatValue, true))
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullFloat)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = fmt.Sprintf("%f", t.Float64)
			}
		},
	}

	NullStringBinder = revel.Binder{
		Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
			if strings.TrimSpace(val) == "" {
				return reflect.ValueOf(NewNullString("", false))
			}
			return reflect.ValueOf(NewNullString(val, true))
		}),
		Unbind: func(output map[string]string, name string, val interface{}) {
			t := val.(NullString)
			if t.Valid == false {
				output[name] = ""
			} else {
				output[name] = t.String
			}
		},
	}
)
