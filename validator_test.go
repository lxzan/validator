package validator

import (
	"testing"
)

type InputsForm struct {
	Name string `format:"required|minLength:3"`
	Age  int64  `format:"min:18"`
	Sex  string `format:"required|switch:male,female"`
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
