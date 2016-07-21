package core

import (
    "time"
    "reflect"
    "errors"
    "github.com/revel/revel"
    "strings"
)

var (
    DateBinder = revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            r, err := time.ParseInLocation("2006-01-02", val, time.Local);
            if err == nil {
                return reflect.ValueOf(Date(r))
            }
            parseError := errors.New(val + "不是有效的日期格式！")
            if strings.Trim(val, " ") == "" {
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
            r, err := time.ParseInLocation("2006-01-02 15:04", val, time.Local);
            if err == nil {
                return reflect.ValueOf(Time(r))
            }
            parseError := errors.New(val + "不是有效的时间格式！")
            if strings.Trim(val, " ") == "" {
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
            r, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local);
            if err == nil {
                return reflect.ValueOf(LongTime(r))
            }
            parseError := errors.New(val + "不是有效的时间格式！")
            if strings.Trim(val, " ") == "" {
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
            if strings.Trim(val, " ") == "" {
                return reflect.ValueOf(NewNullDate(time.Time{}, false))
            }
            r, err := time.ParseInLocation("2006-01-02", val, time.Local);
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
            if strings.Trim(val, " ") == "" {
                return reflect.ValueOf(NewNullTime(time.Time{}, false))
            }
            r, err := time.ParseInLocation("2006-01-02 15:04", val, time.Local);
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
            if strings.Trim(val, " ") == "" {
                return reflect.ValueOf(NewNullLongTime(time.Time{}, false))
            }
            r, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local);
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
)