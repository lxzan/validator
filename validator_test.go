package validator

import (
	"testing"
)

type InputsForm struct {
	Name string `format:"required"`
}

func TestCheck(t *testing.T) {
	LoadLang("en_US", "./data/en_US.ini")
	var inputs = InputsForm{
		Name: "",
	}
	e := Check(&inputs, "en_US")
	println(e)
}
