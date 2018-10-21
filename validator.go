package validator

import (
	"github.com/go-ini/ini"
	"reflect"
	"strings"
)

type CheckFunc func(lang string, key string, val interface{}, limit ...float64) *Error

type Checker map[string]CheckFunc

type Error struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

var checker = Checker{
	"required": func(lang string, key string, val interface{}, limit ...float64) *Error {
		v, ok := val.(string)
		if !ok {
			panic("required only support string type")
		}

		if strings.TrimSpace(v) == "" {
			tpl := GetParam(dict[lang], "tpl", "check_param") + " " + GetParam(dict[lang], "tpl", "not_empty")
			msg := Template(tpl, Form{
				"name": GetParam(dict[lang], "dict", key),
			})
			return &Error{
				Code: 1,
				Msg:  msg,
			}
		}
		return nil
	},
}

var dict = make(map[string]*ini.File)

// loading dictionary
func LoadLang(lang string, file string) {
	f, err := ini.Load(file)
	if err != nil {
		panic("load " + file + " failed")
	}
	dict[lang] = f
}

var defaultLang = "zh_CN"

func Check(inputs interface{}, lang ...string) *Error {
	if len(lang) == 0 {
		lang = append(lang, defaultLang)
	}

	t := reflect.TypeOf(inputs).Elem()
	v := reflect.ValueOf(inputs).Elem()
	for i := 0; i < t.NumField(); i++ {
		rule := t.Field(i).Tag.Get("format")
		val := v.Field(i).Interface()

		name := ToCamel(t.Field(i).Name)
		fields := strings.Split(rule, "|")
		for _, field := range fields {
			f, ok := checker[field]
			if !ok {
				panic(field + "not defined")
			}
			e := f(lang[0], name, val)
			if e != nil {
				return e
			}
		}
	}
	return nil
}
