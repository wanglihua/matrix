package core

import (
    "time"
    "reflect"
    "errors"
    "github.com/revel/revel"
)

var (
    DateBinder = revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            var r time.Time
            var err error
            for _, f := range revel.TimeFormats {
                r, err = time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(Date(r))
                }
            }
            parseError := errors.New(val + "不是有效的日期格式！")
            HandleError(parseError)
            return reflect.Zero(typ)
        }),
        Unbind: func(output map[string]string, name string, val interface{}) {
            var (
                t = val.(time.Time)
                format = revel.DateTimeFormat
                h, m, s = t.Clock()
            )
            if h == 0 && m == 0 && s == 0 {
                format = revel.DateFormat
            }
            output[name] = t.Format(format)
        },
    }

    TimeBinder = revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            var r time.Time
            var err error
            for _, f := range revel.TimeFormats {
                r, err = time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(Time(r))
                }
            }
            parseError := errors.New(val + "不是有效的时间格式！")
            HandleError(parseError)
            return reflect.Zero(typ)
        }),
        Unbind: func(output map[string]string, name string, val interface{}) {
            var (
                t = val.(time.Time)
                format = revel.DateTimeFormat
                h, m, s = t.Clock()
            )
            if h == 0 && m == 0 && s == 0 {
                format = revel.DateFormat
            }
            output[name] = t.Format(format)
        },
    }

    LongTimeBinder = revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            var r time.Time
            var err error
            for _, f := range revel.TimeFormats {
                r, err = time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(LongTime(r))
                }
            }
            parseError := errors.New(val + "不是有效的时间格式！")
            HandleError(parseError)
            return reflect.Zero(typ)
        }),
        Unbind: func(output map[string]string, name string, val interface{}) {
            var (
                t = val.(time.Time)
                format = revel.DateTimeFormat
                h, m, s = t.Clock()
            )
            if h == 0 && m == 0 && s == 0 {
                format = revel.DateFormat
            }
            output[name] = t.Format(format)
        },
    }
)