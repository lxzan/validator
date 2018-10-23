package validator

import (
	"reflect"
	"strconv"
	"strings"
)

var formatMapping = map[string]func(s string) bool{
	"required": IsRequired,
	"email":    IsEmail,
	"ip":       IsIP,
	"url":      IsURL,
	"numeric":  IsNumeric,
}

var valueMapping = map[string]func(s string, value interface{}) (bool, error){
	"min":       min,
	"max":       max,
	"minLength": minLength,
	"maxLength": maxLength,
}

func AddFormatChecker(key string, fn func(s string) bool) {
	formatMapping[key] = fn
}

func isValid(checker *Checker, name string, rule string, value interface{}) *Error {
	arr := strings.Split(rule, ":")
	if arr[0] == "switch" {
		v, ok1 := value.(string)
		if !ok1 {
			panic("switch only support string type")
		}
		ok2, err := testSwitch(rule, v)
		if err != nil {
			panic(name + err.Error())
		}
		if !ok2 {
			msg, err2 := getParam(checker.Dict, "tpl", "switch")
			if err2 != nil {
				panic(name + err2.Error())
			}
			attr, _ := getParam(checker.Dict, "dict", name)
			msg = strings.Replace(msg, ":attr", attr, 1)
			msg = strings.Replace(msg, ":avai", arr[1], 1)
			return &Error{
				Code: 1,
				Msg:  msg,
			}
		}
	} else if len(arr) == 1 {
		v, ok := value.(string)
		if !ok {
			panic(name + " must be string type")
		}
		fn, ok := formatMapping[arr[0]]
		if !ok {
			panic(arr[0] + "'s handler not exist")
		}
		if !fn(v) {
			return &Error{
				Code: 0,
				Msg:  checker.GetMessage(arr[0], name),
			}
		}
	} else {
		limit, err1 := strconv.ParseFloat(arr[1], 64)
		if err1 != nil {
			panic(name + " tag error")
		}
		fn, ok := valueMapping[arr[0]]
		if !ok {
			panic(arr[0] + "'s handler not exist")
		}

		pass, err2 := fn(rule, value)
		if err2 != nil {
			panic(name + err2.Error())
		}

		if !pass {
			return &Error{
				Code: 0,
				Msg:  checker.GetMessage(arr[0], name, limit),
			}
		}
	}
	return nil
}

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
		rule := t.Field(i).Tag.Get("valid")
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
