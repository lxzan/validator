package validator

import (
	"fmt"
	"github.com/go-ini/ini"
	"regexp"
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
	tplValue, e1 := getParam(this.Dict, "tpl", tpl)
	if e1 != nil {
		panic(e1.Error())
	}
	attrValue, e2 := getParam(this.Dict, "dict", attr)
	if e2 != nil {
		panic(e2.Error())
	}

	msg := strings.Replace(tplValue, ":attr", attrValue, 1)
	if len(limit) > 0 {
		l := fmt.Sprintf("%.2f", limit[0])
		re := regexp.MustCompile(`[\.0]+$`)
		l = re.ReplaceAllString(l, "")
		msg = strings.Replace(msg, ":limit", l, 1)
	}
	return msg
}
