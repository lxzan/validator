package validator

import (
	"testing"
)

type InputsForm struct {
	Name string `format:"required"`
}

func TestCheck(t *testing.T) {
	LoadLang("zh_CN", "./data/zh_CN.ini")
	var inputs = InputsForm{
		Name: "",
	}
	e := Check(&inputs, "zh_CN")
	println(e)
}
