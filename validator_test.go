package validator

import (
	"testing"
)

type inputsForm struct {
	Name     string   `valid:"required|minLength:3" default:"caster"`
	Age      int64    `valid:"min:18" default:"1"`
	Sex      string   `valid:"required|switch:male,female"`
	ThreadId []string `valid:"minSize:1"`
}

func TestCheck(t *testing.T) {
	URL := "https://blog/SunWuKong_Hadoop/article/details/74489202?utm_source=blogxgwz0"
	f := IsURL(URL)
	println(f)
}
