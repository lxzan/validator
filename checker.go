package validator

import (
	"fmt"
	"github.com/go-ini/ini"
	"strings"
)

type Checker struct {
	Lang string
	Dict *ini.File
}

func NewChecker(lang string, dict *ini.File) *Checker {
	obj := &Checker{}
	obj.Lang = lang
	obj.Dict = dict
	return obj
}

func (this *Checker) GetMessage(tpl string, attr string, limit ...float64) string {
	tplValue := GetParam(this.Dict, "tpl", tpl)
	attrValue := GetParam(this.Dict, "dict", attr)
	msg := strings.Replace(tplValue, ":attr", attrValue, 1)
	if len(limit) > 0 {
		l := fmt.Sprintf("%f", limit[0])
		msg = strings.Replace(msg, ":limit", l, 1)
	}
	return msg
}
