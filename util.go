package validator

import (
	"strings"
)

type Form map[string]string

func ToCamel(s string) string {
	b := []byte(s)
	if b[0] >= 'A' && b[0] <= 'Z' {
		b[0] += 32
	}
	return string(b)
}

func Template(tpl string, bind Form) string {
	for k, v := range bind {
		tmp := "{" + k + "}"
		tpl = strings.Replace(tpl, tmp, v, -1)
	}
	return tpl
}
