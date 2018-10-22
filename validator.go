package validator

import (
	"reflect"
	"strings"
)

type CheckFunc func(checker *Checker, key string, val interface{}, limit ...float64) *Error

type Error struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var checkFuncs = map[string]CheckFunc{
	"required": func(checker *Checker, key string, val interface{}, limit ...float64) *Error {
		v, ok := val.(string)
		if !ok {
			panic("required only support string type")
		}

		if strings.TrimSpace(v) == "" {
			return &Error{
				Code: 1,
				Msg:  checker.GetMessage("not_empty", "name"),
			}
		}
		return nil
	},

	"email": func(checker *Checker, key string, val interface{}, limit ...float64) *Error {
		v, ok := val.(string)
		if !ok {
			panic("required only support string type")
		}

		if strings.TrimSpace(v) == "" {
			return &Error{
				Code: 1,
				Msg:  checker.GetMessage("not_empty", "name"),
			}
		}
		return nil
	},
}

func Check(inputs interface{}, lang ...string) *Error {
	if len(lang) == 0 {
		lang = append(lang, defaultLang)
	}
	checker := NewChecker(lang[0], dict[lang[0]])

	t := reflect.TypeOf(inputs).Elem()
	v := reflect.ValueOf(inputs).Elem()
	for i := 0; i < t.NumField(); i++ {
		rule := t.Field(i).Tag.Get("format")
		val := v.Field(i).Interface()

		name := ToCamel(t.Field(i).Name)
		fields := strings.Split(rule, "|")
		for _, field := range fields {
			f, ok := checkFuncs[field]
			if !ok {
				panic(field + "not defined")
			}
			e := f(checker, name, val)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
