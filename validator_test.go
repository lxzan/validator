package validator

import (
	"testing"
)

type InputsForm struct {
	Name string `valid:"required|minLength:3"`
	Age  int64  `valid:"min:18"`
	Sex  string `valid:"required|switch:male,female"`
}

func TestCheck(t *testing.T) {
	LoadLang("zh_CN", "./data/zh_CN.ini")
	var inputs = InputsForm{
		Name: "lxz",
		Age:  19,
		Sex:  "unknown",
	}

	e := Check(&inputs, "zh_CN")
	println(e)
}
