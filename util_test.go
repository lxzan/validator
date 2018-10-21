package validator

import "testing"

func TestToCamel(t *testing.T) {
	s := ToCamel("TestMe")
	println(s)
}
