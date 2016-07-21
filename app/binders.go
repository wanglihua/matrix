package app

import (
    "reflect"
    "github.com/revel/revel"
    "matrix/core"
)

func registerBinders() {
    revel.TypeBinders[reflect.TypeOf(core.Date{})] = core.DateBinder
    revel.TypeBinders[reflect.TypeOf(core.Time{})] = core.TimeBinder
    revel.TypeBinders[reflect.TypeOf(core.LongTime{})] = core.LongTimeBinder
    revel.TypeBinders[reflect.TypeOf(core.NullDate{})] = core.NullDateBinder
    revel.TypeBinders[reflect.TypeOf(core.NullTime{})] = core.NullTimeBinder
    revel.TypeBinders[reflect.TypeOf(core.NullLongTime{})] = core.NullLongTimeBinder
}
