package validator

import (
	"reflect"
	"strings"
)

type CheckFunc func(key string, val interface{}, limit ...float64) *Error

type Checker map[string]CheckFunc

type Error struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var checker = Checker{
	"required": func(key string, val interface{}, limit ...float64) *Error {
		v, ok := val.(string)
		if !ok {
			panic("required only support string type")
		}

		if strings.TrimSpace(v) == "" {
			return &Error{
				Code: 1,
				Msg:  key + "不能为空",
			}
		}
		return nil
	},
}

func Check(inputs interface{}) *Error {
	t := reflect.TypeOf(inputs).Elem()
	v := reflect.ValueOf(inputs).Elem()
	for i := 0; i < t.NumField(); i++ {
		rule := t.Field(i).Tag.Get("format")
		val := v.Field(i).Interface()

		name := t.Field(i).Name
		fields := strings.Split(rule, "|")
		for _, field := range fields {
			f, ok := checker[field]
			if !ok {
				panic(field + "未定义")
			}
			e := f(name, val)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
