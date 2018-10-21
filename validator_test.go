package validator

import (
	"testing"
)

type Form struct {
	Name string `format:"required"`
}

func TestCheck(t *testing.T) {
	var inputs = Form{
		Name: "",
	}
	e := Check(&inputs)
	println(e)
}
