package validator

import "github.com/pkg/errors"

type Form map[string]string

func ToCamel(s string) string {
	b := []byte(s)
	if b[0] >= 'A' && b[0] <= 'Z' {
		b[0] += 32
	}
	return string(b)
}

func ToFloat64(value interface{}) (float64, error) {
	v1, ok := value.(int)
	if ok {
		return float64(v1), nil
	}

	v2, ok := value.(int64)
	if ok {
		return float64(v2), nil
	}

	v3, ok := value.(float32)
	if ok {
		return float64(v3), nil
	}

	v4, ok := value.(float64)
	if ok {
		return float64(v4), nil
	}

	return 0, errors.New(" only support int, int64, float32, float64")
}
