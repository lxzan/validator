package validator

import (
	"reflect"
	"strings"
)

type Error struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
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
		if rule == "" {
			continue
		}
		val := v.Field(i).Interface()

		name := ToCamel(t.Field(i).Name)
		fields := strings.Split(rule, "|")
		for _, field := range fields {
			e := isValid(checker, name, field, val)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
