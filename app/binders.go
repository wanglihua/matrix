package app

import (
    "time"
    "reflect"
    "github.com/revel/revel"
    "matrix/core"
)

func registerBinders() {

    //core.Date
    dateBinder := revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            for _, f := range revel.TimeFormats {
                 r, err := time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(core.Date(r))
                } else {
                    core.HandleError(err)
                }
            }
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

    revel.TypeBinders[reflect.TypeOf(core.Date{})] = dateBinder

    //core.Time
    timeBinder := revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            for _, f := range revel.TimeFormats {
                 r, err := time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(core.Time(r))
                } else {
                    core.HandleError(err)
                }
            }
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

    revel.TypeBinders[reflect.TypeOf(core.Time{})] = timeBinder

    //core.LongTime
    longTimeBinder := revel.Binder{
        Bind: revel.ValueBinder(func(val string, typ reflect.Type) reflect.Value {
            for _, f := range revel.TimeFormats {
                 r, err := time.Parse(f, val);
                if err == nil {
                    return reflect.ValueOf(core.LongTime(r))
                } else {
                    core.HandleError(err)
                }
            }
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

    revel.TypeBinders[reflect.TypeOf(core.LongTime{})] = longTimeBinder
}
