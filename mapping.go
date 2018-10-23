package validator

import (
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
